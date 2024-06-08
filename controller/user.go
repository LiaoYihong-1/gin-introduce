package controller

import (
	"gin/request"
	"gin/response"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserById(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	status := http.StatusOK
	user := response.User{}
	if e != nil {
		status = http.StatusBadRequest
	} else {
		u, er := service.FindUserById(id)
		user = utils.ModelToResponse(u)
		if er != nil {
			status = http.StatusNotFound
			user = response.User{}
		}
		if u.ID == 0 {
			status = http.StatusNotFound
		}
	}
	if status != http.StatusOK {
		c.Status(status)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
func CreateUser(c *gin.Context) {
	userRequest := request.User{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	u := utils.UserRequestToModel(userRequest)
	service.CreateUser(u)
	c.Status(http.StatusNoContent)
}
func GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	response := make([]response.User, len(users))
	for i, user := range users {
		response[i] = utils.ModelToResponse(user)
	}
	c.JSON(http.StatusOK, response)
}
func DeleteUser(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	status := http.StatusNoContent
	if e != nil {
		status = http.StatusBadRequest
	} else {
		err := service.DeleteUser(id)
		if err != nil {
			status = http.StatusNotFound
		}
	}
	c.Status(status)
}
func UpdateUser(c *gin.Context) {
	id, pathError := strconv.Atoi(c.Param("id"))
	requestUser := request.User{}
	bodyError := c.BindJSON(&requestUser)
	status := http.StatusNoContent
	if pathError != nil || bodyError != nil {
		status = http.StatusBadRequest
	} else {
		user := utils.UserRequestToModel(requestUser)
		user.ID = id
		err := service.UpdateUser(user)
		if err != nil {
			status = http.StatusNotFound
		}
	}
	c.Status(status)
}
func Register(c *gin.Context) {
	userRequest := request.User{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	u := utils.UserRequestToModel(userRequest)
	if service.Register(u) != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusNoContent)
}
func Login(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	token, err := service.Login(email, password)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
