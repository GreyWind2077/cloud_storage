package file

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)
	//获取图像类型文件
	imgFiles := service.GetUserFilesByType(2, user.FileStoreId)

	c.HTML(http.StatusOK, "image-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"imgFiles":      imgFiles,
		"imgCount":      len(imgFiles),
		"currImg":       "active",
		"currClass":     "active",
	})
}
