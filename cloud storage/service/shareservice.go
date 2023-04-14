package service

import (
	"cloud_storage/dao"
	"cloud_storage/models"
	"cloud_storage/utils"
	"strings"
	"time"
)

func CreateShare(code, username string, fid int) string {
	share := models.Share{
		Code:     strings.ToLower(code),
		Id:       fid,
		Username: username,
		Hash:     utils.Md5Crypt(code, string(time.Now().Unix())), //拿当前时间当加密盐
	}
	dao.CreateShare(&share)
	return share.Hash
}

func GetShare(fhash string) (share models.Share) {

	return dao.FindShareByHash(fhash)

}

// 验证提取码
func VerifyShareCode(fid, code string) bool {
	return dao.FindShareCode(fid, code)
}
