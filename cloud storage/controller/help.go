package controller

import (
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Help(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	//获取用户文件使用明细数量
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)

	c.HTML(http.StatusOK, "help.html", gin.H{
		"currHelp":      "active",
		"user":          user,
		"fileDetailUse": fileDetailUse,
	})
}
