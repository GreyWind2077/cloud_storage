package controller

import (
	"cloud_storage/dao/oss"
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"github.com/lifei6671/gocaptcha"
	"net/http"
	"strconv"
	"strings"
)

// 创建分享文件
func ShareFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := service.FindUser(openId)

	fId := c.Query("id")
	url := c.Query("url")
	//生成验证码
	code := gocaptcha.RandText(4)

	fileId, _ := strconv.Atoi(fId)
	hash := service.CreateShare(code, user.UserName, fileId)

	c.JSON(http.StatusOK, gin.H{
		"url":  url + "?f=" + hash,
		"code": code,
	})
}

// 分享文件页面
func SharePass(c *gin.Context) {
	f := c.Query("f")

	//获取分享信息
	shareInfo := service.GetShare(f)
	//获取文件信息
	file := service.GetFileById(strconv.Itoa(shareInfo.FileId))

	c.HTML(http.StatusOK, "share.html", gin.H{
		"id":       shareInfo.FileId,
		"username": shareInfo.Username,
		"fileType": file.Type,
		"filename": file.FileName + file.PostFix,
		"hash":     shareInfo.Hash,
	})
}

// 下载分享文件
func DownloadShareFile(c *gin.Context) {
	fileId := c.Query("id")
	code := c.Query("code")
	hash := c.Query("hash")

	fileInfo := service.GetFileById(fileId)

	//校验提取码
	if ok := service.VerifyShareCode(fileId, strings.ToLower(code)); !ok {
		c.Redirect(http.StatusMovedPermanently, "/file/share?f="+hash)
		return
	}

	//从oss获取文件
	fileData := oss.DownLoadFromOSS(fileInfo.FileHash, fileInfo.PostFix)
	//下载次数+1
	service.UpdateLoadNum(fileId)

	c.Header("Content-disposition", "attachment;filename=\""+fileInfo.FileName+fileInfo.PostFix+"\"")
	c.Data(http.StatusOK, "application/octect-stream", fileData)
}
