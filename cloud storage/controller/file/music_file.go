package file

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MusicFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)
	//获取音频类型文件
	musicFiles := service.GetUserFilesByType(4, user.FileStoreId)

	c.HTML(http.StatusOK, "music-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"musicFiles":    musicFiles,
		"musicCount":    len(musicFiles),
		"currMusic":     "active",
		"currClass":     "active",
	})
}
