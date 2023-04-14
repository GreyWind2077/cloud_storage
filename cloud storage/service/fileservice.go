package service

import (
	"cloud_storage/dao"
	"cloud_storage/models"
	"cloud_storage/utils"
	"path"
	"strconv"
	"strings"
	"time"
)

// 添加文件数据
func CreateFile(filename, filehash, fileid string, filesize int64, filestoreid int) {
	var size string
	fileSuffix := path.Ext(filename)

	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	fid, _ := strconv.Atoi(fileid)

	if filesize < 1048576 {
		size = strconv.FormatInt(filesize/1024, 10) + "KB"
	} else {
		size = strconv.FormatInt(filesize/102400, 10) + "MB"
	}

	myFile := models.MyFile{
		FileName:       filePrefix,
		FileHash:       filehash,
		FileStoreId:    filestoreid,
		FilePath:       "",
		DownLoadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fid,
		Size:           filesize / 1024,
		SizeStr:        size,
		Type:           utils.GetFileType(fileSuffix),
		PostFix:        strings.ToLower(fileSuffix),
	}
	dao.CreateMyFile(&myFile)

}

// 获取用户文件
func GetUserFiles(parentid string, storeid int) (files []models.MyFile) {

	dao.FindUserFiles(&files, parentid, storeid)

	return files
}

// 根据文件类型获取用户文件
func GetUserFilesByType(filetype, filestoreid int) (files []models.MyFile) {
	dao.FindUserFilesByType(&files, filetype, filestoreid)
	return files
}

// 获取用户文件数量
func GetUserFilesCount(storeid int) (filecount int) {
	return dao.GetUserFilesCount(storeid)
}

// 根据id查找文件
func GetFileById(fid string) (file models.MyFile) {
	dao.FindFileById(&file, fid)
	return file
}

// 删除文件
func DeleteUserFile(fid, folderid string, storeid int) {
	dao.DeleteFile(fid, folderid, storeid)

}

// 更新用户文件容量
func UpdateSize(size int64, filestoreid int) {
	dao.UpdateSize(size, filestoreid)
}

// 更新下载量
func UpdateLoadNum(fid string) {
	dao.UpdateLoadNum(fid)
}

// 获取用户网盘使用情况
func GetFileStoreUse(filestoreid int) map[string]int64 {
	return dao.GetFileStoreUse(filestoreid)
}

func IsTheSameName(fid, filename string) bool {
	return dao.FindTheSameName(fid, filename)
}

func IsInOss(fileHash string) bool {
	return dao.IsInOSS(fileHash)
}
