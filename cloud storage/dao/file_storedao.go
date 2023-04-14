package dao

import (
	"cloud_storage/dao/database"
	"cloud_storage/models"
)

func CreateFileStore(filestore *models.FileStore) {
	database.DB.Create(filestore)
}
func FindUserStoreById(userid int) (filestore models.FileStore) {
	database.DB.Find(&filestore, "user_id = ?", userid)
	return filestore
}

func IsEnough(filesize int64, filestoreid int) bool {
	var filestore models.FileStore
	database.DB.Find(&filestore, filestoreid)

	if filestore.MaxSize-(filesize/1024) < 0 {
		return false
	}
	return true
}
