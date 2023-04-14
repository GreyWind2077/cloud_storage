package dao

import (
	"cloud_storage/dao/database"
	"cloud_storage/models"
	"strconv"
)

func CreateFileFolder(filefolder *models.FileFolder) {
	database.DB.Create(filefolder)
}

func FindParentFolderById(fid string) (file models.FileFolder) {
	database.DB.Find(&file, "id=?", fid)
	return file
}

func FindFileFolder(parentid string, filestoreid int) (filefolders []models.FileFolder) {
	database.DB.Order("time desc").Find(&filefolders, "parent_folder_id=? and file_store_id=?", parentid, filestoreid)
	return filefolders

}

func FindFileFolderById(fid string) (filefolder models.FileFolder) {
	database.DB.Find(&filefolder, "id=?", fid)
	return filefolder
}

func FindFileFolderParent(folder models.FileFolder, folders []models.FileFolder) []models.FileFolder {
	var parentFolder models.FileFolder

	if folder.ParentFolderId != 0 {
		database.DB.Find(&parentFolder, "id=?", folder.ParentFolderId)
		folders = append(folders, parentFolder)

		return FindFileFolderParent(parentFolder, folders)
	}

	//反转
	for i, j := 0, len(folders)-1; i < j; i, j = i+1, j-1 {
		folders[i], folders[j] = folders[j], folders[i]
	}
	return folders
}

func UserFileFolderCount(filestoreid int) (fildfoldercount int) {
	var fileFolder []models.FileFolder
	database.DB.Find(&fileFolder, "file_store_id", filestoreid).Count(&fildfoldercount)
	return fildfoldercount
}

func DeleteFileFolder(fid string) bool {
	database.DB.Begin()
	var filefolder1, filefolder2 models.FileFolder
	database.DB.Where("id=?", fid).Delete(models.FileFolder{})
	database.DB.Where("parent_folder_id=?", fid).Delete(models.MyFile{})
	database.DB.Find(&filefolder1, "parent_folder_id=?", fid)
	database.DB.Where("parent_folder_id=?", fid).Delete(models.FileFolder{})
	database.DB.Find(&filefolder2, "parent_folder_id", filefolder1.Id)
	if filefolder2.Id != 0 {
		return DeleteFileFolder(strconv.Itoa(filefolder1.Id))
	}

	database.DB.Commit()

	return true

}

func UpdateFolderName(fid, fname string) {
	var filefolder models.FileFolder
	database.DB.Model(&filefolder).Where("id=?", fid).Update("file_folder_name", fname)
}
