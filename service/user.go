package service

import (
	"gin/database"
	"gin/exception"
	"gin/models"
	"gin/repository"
	"gin/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func FindUserById(id int) (models.User, error) {
	if id > 0 {
		return repository.GetUser(id), nil
	}
	return models.User{}, &exception.UserNotFoundError{}
}
func CreateUser(user models.User) {
	repository.CreateUser(user)
}
func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}
func DeleteUser(id int) error {
	if id > 0 {
		user := repository.GetUser(id)
		if user.ID == 0 {
			return &exception.UserNotFoundError{}
		}
		repository.DeleteUser(user)
	} else {
		return &exception.UserNotFoundError{}
	}
	return nil
}
func UpdateUser(user models.User) error {
	if user.ID > 0 {
		if repository.GetUser(user.ID).ID == 0 {
			return &exception.UserNotFoundError{}
		}
		repository.UpdateUser(user)
		return nil
	}
	return &exception.UserNotFoundError{}
}
func Login(username string, password string) (string, error) {
	user := repository.FindUserByEmail(username)
	if user.ID == 0 {
		return "", &exception.UserNotFoundError{}
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", &exception.BadRequestError{}
	}
	token, err := utils.GenToken(user.ID)
	if err != nil {
		return "", err
	}
	if !repository.IsLoadedUser(user.ID) {
		repository.SaveUserPrivilege(user.ID, user.Privileges)
	}
	return token, nil
}
func Register(user models.User) error {
	hashPassword, hashError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashError != nil {
		return hashError
	}
	user.Password = string(hashPassword)
	CreateUser(user)
	return nil
}
func GetUserPrivilegesInRedis(id int) ([]utils.PrivilegeType, error) {
	if !repository.IsLoadedUser(id) {
		return nil, &exception.UserNotFoundError{Info: "User not login"}
	}
	key := "user:" + strconv.Itoa(id)
	privileges, _ := database.Redis.LRange(key, 0, -1).Result()
	var result []utils.PrivilegeType
	for _, v := range privileges {
		result = append(result, utils.PrivilegeType(v))
	}
	return result, nil
}
