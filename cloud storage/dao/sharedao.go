package dao

import (
	"cloud_storage/dao/database"
	"cloud_storage/models"
)

func CreateShare(share *models.Share) {
	database.DB.Create(share)
}

func FindShareByHash(fhash string) (share models.Share) {
	database.DB.Find(&share, "hash=?", fhash)

	return share
}

func FindShareCode(fid, code string) bool {
	var share models.Share
	database.DB.Find(&share, "file_id=? and code =?", fid, code)

	if share.Id == 0 {
		return false
	}

	return true
}
