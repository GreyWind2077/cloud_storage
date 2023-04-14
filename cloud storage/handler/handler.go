package handler

import (
	"cloud_storage/config"
	"cloud_storage/dao/oss"
	"cloud_storage/service"
	"cloud_storage/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func LoginHandler(c *gin.Context) {
	state := "xxxxxxx"
	url := "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + config.Cfg.QQ.App_id + "&redirect_uri=" + config.Cfg.QQ.Redirect_url + "&state=" + state
	c.Redirect(http.StatusMovedPermanently, url)
}

// 处理上传文件
func UploadHandler(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := service.FindUser(openId)

	Fid := c.GetHeader("id")
	//接收上传文件
	file, head, err := c.Request.FormFile("file")

	//判断当前文件夹是否有同名文件
	if ok := service.IsTheSameName(Fid, head.Filename); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
		})
		return
	}

	//判断用户的容量是否足够
	if ok := service.GetCapacity(head.Size, user.FileStoreId); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 503,
		})
		return
	}

	if err != nil {
		log.Println("文件上传错误", err.Error())
		return
	}
	defer file.Close()

	//文件保存本地的路径
	location := config.Cfg.App.Location + head.Filename

	//在本地创建一个新的文件
	newFile, err := os.Create(location)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("文件拷贝错误", err.Error())
		return
	}

	//将光标移至开头
	_, _ = newFile.Seek(0, 0)
	fileHash := utils.GetSHA256HashCode(newFile)

	//通过hash判断文件是否已上传过oss
	if ok := service.IsInOss(fileHash); ok {
		//上传至阿里云oss
		go oss.UpLoadToOSS(head.Filename, fileHash)
	}
	//新建文件信息
	service.CreateFile(head.Filename, fileHash, Fid, fileSize, user.FileStoreId)
	//上传成功减去相应剩余容量
	service.UpdateSize(fileSize/1024, user.FileStoreId)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
