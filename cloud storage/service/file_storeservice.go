package service

import (
	"cloud_storage/dao"
	"cloud_storage/models"
)

// 根据id查询仓库信息
func GetUserFileStore(userid int) (filestore models.FileStore) {
	return dao.FindUserStoreById(userid)
}

// 判断用户容量是否充足
func GetCapacity(fileSize int64, filestoreid int) bool {
	return dao.IsEnough(fileSize, filestoreid)
}
