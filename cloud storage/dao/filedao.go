package dao

import (
	"cloud_storage/dao/database"
	"cloud_storage/models"
	"path"
	"strings"
)

func CreateMyFile(file *models.MyFile) {
	database.DB.Create(file)
}

func GetUserFilesCount(storeid int) (filecount int) {
	var file []models.MyFile
	database.DB.Find(&file, "file_store_id=?", storeid).Count(&filecount)
	return filecount
}

func FindUserFiles(file *[]models.MyFile, parentid string, storeid int) {
	database.DB.Find(file, "file_store_id = ? and parent_folder_id =?", storeid, parentid)
}

func FindUserFilesByType(file *[]models.MyFile, filetype, filestoreid int) {
	database.DB.Find(file, "type=? and file_store_id=?", filetype, filestoreid)

}

func FindFileById(file *models.MyFile, fileid string) {
	database.DB.Find(file, "id = ?", fileid)
}

func DeleteFile(fid, folderid string, storeid int) {
	database.DB.Where("id =? and file_store_id=? and parent_folder_id =?", fid, storeid, folderid).Delete(models.MyFile{})
}

func UpdateSize(size int64, filestoreid int) {
	database.DB.Begin()
	var filestore models.FileStore
	database.DB.First(&filestore, filestoreid)

	filestore.CurrentSize += size / 1024
	filestore.MaxSize -= size / 1024
	database.DB.Save(&filestore)

	database.DB.Commit()
}

func UpdateLoadNum(fid string) {
	var file models.MyFile
	database.DB.First(&file, fid)
	var downloadnum = file.DownLoadNum

	database.DB.Where("where download_num=?", downloadnum).Save(&file)

}

var filetype = []string{
	"docCount",
	"imgCount",
	"videoCount",
	"musicCount",
	"otherCount",
}

func GetFileStoreUse(filestore int) map[string]int64 {
	var files []models.MyFile
	filestoreuse := make(map[string]int64, 0)
	var Count int64
	for i := 0; i < 5; i++ {
		Count = database.DB.Find(&files, "file_store_id =? and type = ?", filestore, i+1).RowsAffected
		filestoreuse[filetype[i]] = Count
	}
	return filestoreuse
}

func IsInOSS(filehash string) bool {
	var file models.MyFile
	database.DB.Find(&file, "file_hash=?", filehash)
	if file.FileHash == "" {
		return false
	}
	return true
}

func FindTheSameName(fid, filename string) bool {
	var file models.MyFile
	fileSuffix := strings.ToLower(path.Ext(filename))
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	database.DB.Find(&file, "parent_folder_id = ? and file_name = ? and postfix = ?", fid, filePrefix, fileSuffix)

	if file.Size > 0 {
		return false
	}
	return true

}
