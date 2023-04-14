package file

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DocFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)
	//获取文档类型文件
	docFiles := service.GetUserFilesByType(1, user.FileStoreId)

	c.HTML(http.StatusOK, "doc-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"docFiles":      docFiles,
		"docCount":      len(docFiles),
		"currDoc":       "active",
		"currClass":     "active",
	})
}
