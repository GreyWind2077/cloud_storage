package file

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)
	//获取视频类型文件
	videoFiles := service.GetUserFilesByType(3, user.FileStoreId)

	c.HTML(http.StatusOK, "video-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"videoFiles":    videoFiles,
		"videoCount":    len(videoFiles),
		"currVideo":     "active",
		"currClass":     "active",
	})
}
