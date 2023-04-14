package dao

import (
	"cloud_storage/dao/database"
	"cloud_storage/models"
)

func CreateUser(user *models.User) {

	database.DB.Create(user)
}

func SaveUser(user *models.User) {
	database.DB.Save(user)
}

func FindUser(user *models.User, openId interface{}) {
	database.DB.Find(user, "open_id=?", openId)
}
