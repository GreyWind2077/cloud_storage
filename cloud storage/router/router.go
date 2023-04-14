package router

import (
	"cloud_storage/controller"
	"cloud_storage/controller/file"
	"cloud_storage/handler"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func LoadRoute() {
	//1.页面 views 2.数据（json) 3.静态资源
	Router = gin.Default()
	Router.GET("/", controller.Login)
	Router.GET("/qq_login", handler.LoginHandler)
	Router.GET("/callbackQQ", controller.QQLogin)
	Router.GET("/file/share", controller.SharePass)
	Router.GET("/file/shareDownload", controller.DownloadShareFile)

	cloud := Router.Group("cloud")
	cloud.Use(controller.CheckLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/files", file.File)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/doc-files", file.DocFile)
		cloud.GET("/image-files", file.ImageFile)
		cloud.GET("/video-files", file.VideoFile)
		cloud.GET("/music-files", file.MusicFile)
		cloud.GET("/other-files", file.OtherFile)
		cloud.GET("/logout", controller.Logout)
		cloud.GET("/downloadFile", file.DownloadFile)
		cloud.GET("/deleteFile", file.DeleteFile)
		cloud.GET("/deleteFolder", file.DeleteFileFolder)
		cloud.GET("/help", controller.Help)
	}

	{
		cloud.POST("/uploadFile", handler.UploadHandler)
		cloud.POST("/addFolder", file.AddFolder)
		cloud.POST("/updateFolder", file.UpdateFileFolder)
		cloud.POST("/getQrCode", controller.ShareFile)
	}
}
