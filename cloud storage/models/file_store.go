package models

type FileStore struct {
	Id          int
	UserId      int
	CurrentSize int64
	MaxSize     int64
}
