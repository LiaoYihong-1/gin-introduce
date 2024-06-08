package repository

import (
	"gin/database"
	"gin/models"
	"strconv"
	"time"
)

func SaveUserPrivilege(id int, privileges []models.Privilege) {
	userKey := "user:" + strconv.Itoa(id)
	for _, privilege := range privileges {
		database.Redis.LPush(userKey, privilege.Name)
	}
	database.Redis.Expire(userKey, time.Minute*2)
}
func IsLoadedUser(id int) bool {
	key := "user:" + strconv.Itoa(id)
	load, err := database.Redis.Exists(key).Result()
	if err != nil {
		return false
	}
	return load == 1
}
