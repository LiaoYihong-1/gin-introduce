package repository

import (
	"gin/database"
	"gin/models"
)

func CreateUser(user models.User) {
	database.DB.Create(&user)
}
func GetUser(id int) models.User {
	user := models.User{}
	database.DB.Preload("Privileges").First(&user, id)
	return user
}
func DeleteUser(user models.User) {
	database.DB.Delete(&user)
}
func UpdateUser(user models.User) {
	database.DB.Save(&user)
}
func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}
func FindUserByEmail(email string) models.User {
	var user models.User
	database.DB.Preload("Privileges").Where("email = ?", email).First(&user)
	return user
}
