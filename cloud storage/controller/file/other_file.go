package file

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OtherFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)
	//获取音频类型文件
	otherFiles := service.GetUserFilesByType(5, user.FileStoreId)

	c.HTML(http.StatusOK, "other-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"otherFiles":    otherFiles,
		"otherCount":    len(otherFiles),
		"currOther":     "active",
		"currClass":     "active",
	})
}
