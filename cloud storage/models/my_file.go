package models

// MyFile文件夹实体类
type MyFile struct {
	Id             int
	FileName       string
	FileHash       string
	FileStoreId    int
	FilePath       string
	DownLoadNum    int
	UploadTime     string
	ParentFolderId int
	Size           int64
	SizeStr        string
	Type           int
	PostFix        string
}
