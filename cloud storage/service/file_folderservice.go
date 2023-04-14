package service

import (
	"cloud_storage/dao"
	"cloud_storage/models"
	"log"
	"strconv"
	"time"
)

func CreateFolder(folderName, parentId string, filestoreid int) {
	parentIdint, err := strconv.Atoi(parentId)
	if err != nil {
		log.Fatal("父类id错误")
	}

	fileFolder := models.FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentIdint,
		FileStoreId:    filestoreid,
		Time:           time.Now().Format("2006-04-02 15:01:05"),
	}
	dao.CreateFileFolder(&fileFolder)
}

// 获取父类的id
func GetParentFolder(fId string) (file models.FileFolder) {

	return dao.FindParentFolderById(fId)
}

// 获取仓库中的文件
func GetAllFileFolder(parentid string, filestoreid int) (filefolder []models.FileFolder) {
	return dao.FindFileFolder(parentid, filestoreid)
}

// 获取当前目录的信息
func GetCurrentFolder(fid string) (filefolder models.FileFolder) {
	return dao.FindFileFolderById(fid)
}

func GetCurrentFolderParent(folder models.FileFolder, folders []models.FileFolder) []models.FileFolder {
	return dao.FindFileFolderParent(folder, folders)
}

func GetUserFileFolderCount(filestoreid int) (fileFolderCount int) {
	return dao.UserFileFolderCount(filestoreid)
}

func DeleteFileFolder(fid string) bool {
	return dao.DeleteFileFolder(fid)
}

func ModifyFolderName(fid, fname string) {
	dao.UpdateFolderName(fid, fname)
}
