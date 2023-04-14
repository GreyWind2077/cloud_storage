package models

type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}
