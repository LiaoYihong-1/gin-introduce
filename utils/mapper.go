package utils

import (
	"gin/models"
	"gin/request"
	"gin/response"
)

func UserRequestToModel(user request.User) models.User {
	u := models.User{}
	u.Name = user.Name
	u.Email = user.Email
	u.Age = user.Age
	u.Password = user.Password
	return u
}

func ModelToResponse(user models.User) response.User {
	u := response.User{}
	u.Name = user.Name
	u.Email = user.Email
	u.Age = user.Age
	return u
}
