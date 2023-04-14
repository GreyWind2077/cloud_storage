package file

import (
	"cloud_storage/dao/oss"
	"cloud_storage/models"
	"cloud_storage/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 全部文件
func File(c *gin.Context) {
	openId, _ := c.Get("openId")
	fid := c.DefaultQuery("fId", "0")
	user := service.FindUser(openId)
	files := service.GetUserFiles(fid, user.FileStoreId)
	fileFolder := service.GetAllFileFolder(fid, user.FileStoreId)
	parentFolder := service.GetParentFolder(fid)
	currentAllParent := service.GetCurrentFolderParent(parentFolder, make([]models.FileFolder, 0))
	currentFolder := service.GetCurrentFolder(fid)
	fileDetailUse := service.GetFileStoreUse(user.FileStoreId)

	c.HTML(http.StatusOK, "files.html", gin.H{
		"currAll":          "active",
		"user":             user,
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"files":            files,
		"fileFolder":       fileFolder,
		"parentFolder":     parentFolder,
		"currentAllParent": currentAllParent,
		"fileDetailUse":    fileDetailUse,
	})

}

func AddFolder(c *gin.Context) {
	openid, _ := c.Get("openId")
	user := service.FindUser(openid)
	folderName := c.PostForm("fileFolderName")
	parentId := c.DefaultPostForm("parentFolderId", "0")

	//新建文件夹数据
	service.CreateFolder(folderName, parentId, user.FileStoreId)

	//获取父文件夹信息
	parent := service.GetParentFolder(parentId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+parentId+"&fName="+parent.FileFolderName)

}

func DownloadFile(c *gin.Context) {
	fId := c.Query("fId")

	file := service.GetFileById(fId)
	if file.FileHash == "" {
		return
	}

	//从oss获取文件
	fileData := oss.DownLoadFromOSS(file.FileHash, file.PostFix)
	//下载次数+1
	service.UpdateLoadNum(fId)

	c.Header("Content-disposition", "attachment;filename=\""+file.FileName+file.PostFix+"\"")
	c.Data(http.StatusOK, "application/octect-stream", fileData)
}

// 删除文件
func DeleteFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := service.FindUser(openId)

	fId := c.DefaultQuery("fId", "")
	folderId := c.Query("folder")
	if fId == "" {
		return
	}

	//删除数据库文件数据
	service.DeleteUserFile(fId, folderId, user.FileStoreId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fid="+folderId)
}

// 删除文件夹
func DeleteFileFolder(c *gin.Context) {
	fId := c.DefaultQuery("fId", "")
	if fId == "" {
		return
	}
	//获取要删除的文件夹信息 取到父级目录重定向
	folderInfo := service.GetCurrentFolder(fId)

	//删除文件夹并删除文件夹中的文件信息
	service.DeleteFileFolder(fId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+strconv.Itoa(folderInfo.ParentFolderId))
}

// 修改文件夹名
func UpdateFileFolder(c *gin.Context) {
	fileFolderName := c.PostForm("fileFolderName")
	fileFolderId := c.PostForm("fileFolderId")

	fileFolder := service.GetCurrentFolder(fileFolderId)

	service.ModifyFolderName(fileFolderId, fileFolderName)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+strconv.Itoa(fileFolder.ParentFolderId))
}
