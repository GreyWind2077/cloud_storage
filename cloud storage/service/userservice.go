package service

import (
	"cloud_storage/dao"
	"cloud_storage/models"
	"time"
)

// 创建用户，同时创建仓库
func CreateUser(openId, username, image string) {

	user := models.User{
		OpenId:       openId,
		FileStoreId:  0,
		UserName:     username,
		RegisterTime: time.Now(),
		ImagePath:    image,
	}
	dao.CreateUser(&user)

	filestore := models.FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048578, //每个用户的存储上限为1G,即104878kb
	}
	dao.CreateFileStore(&filestore)
	user.FileStoreId = filestore.Id
	dao.SaveUser(&user)

}

func FindUser(openId interface{}) models.User {
	var user models.User
	dao.FindUser(&user, openId)
	return user
}
