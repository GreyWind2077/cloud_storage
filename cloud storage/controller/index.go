package controller

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	openid, _ := c.Get("openId")
	user := service.FindUser(openid)
	userfilestore := service.GetUserFileStore(user.Id)
	fileCount := service.GetUserFilesCount(user.FileStoreId)
	filefoldercount := service.GetUserFileFolderCount(user.FileStoreId)
	fileUse := service.GetFileStoreUse(user.FileStoreId)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":            user,
		"currIndex":       "active",
		"userFileStore":   userfilestore,
		"fileCount":       fileCount,
		"fileFolderCount": filefoldercount,
		"fileDetailUse":   fileUse,
	})
}
