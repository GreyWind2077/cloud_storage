package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// 文档类型
var document = []string{
	".doc",
	".doxc",
	".txt",
	".pdf",
}

// 图片类型
var photo = []string{
	".jpg",
	".png",
	".gif",
	".jpeg",
}

var video = []string{
	".mp4",
	".avi",
	".mov",
	"rmvb",
	".rm",
}
var music = []string{
	".mp3",
	".cda",
	".wav",
	".wma",
	".ogg",
}

// 判断是否是该种类型
func InArray(str string, strlist []string) bool {
	for _, s := range strlist {
		if s == str {
			return true
		}
	}

	return false
}

func GetFileType(fileSuffix string) int {
	fileSuffix = strings.ToLower(fileSuffix)
	if InArray(fileSuffix, document) {
		return 1
	}
	if InArray(fileSuffix, photo) {
		return 2
	}
	if InArray(fileSuffix, video) {
		return 3
	}

	if InArray(fileSuffix, music) {
		return 4

	}

	return 5

}

// MD5加密
func Md5Crypt(str string, salt ...interface{}) string {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))

}

// SHA256生成哈希值,加密
func GetSHA256HashCode(file *os.File) string {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	_, _ = io.Copy(hash, file)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

}
