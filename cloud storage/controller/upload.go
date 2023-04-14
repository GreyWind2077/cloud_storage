package controller

import (
	"cloud_storage/models"
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 上传文件页面
func Upload(c *gin.Context) {
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	//获取用户信息
	user := service.FindUser(openId)
	//获取当前目录信息
	currentFolder := service.GetCurrentFolder(fId)
	//获取当前目录所有的文件夹信息
	fileFolders := service.GetAllFileFolder(fId, user.FileStoreId)
	//获取父级的文件夹信息
	parentFolder := service.GetParentFolder(fId)
	//获取当前目录所有父级
	currentAllParent := service.GetCurrentFolderParent(parentFolder, make([]models.FileFolder, 0))
	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"user":             user,
		"currUpload":       "active",
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"fileFolders":      fileFolders,
		"parentFolder":     parentFolder,
		"currentAllParent": currentAllParent,
		"fileDetailUse":    fileDetailUse,
	})
}
