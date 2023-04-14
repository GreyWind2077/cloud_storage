package oss

import (
	"cloud_storage/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"log"
	"path"
)

func UpLoadToOSS(filename, filehash string) {
	fileSuffix := path.Ext(filename)

	// 创建OSSClient实例。
	client, err := oss.New(config.Cfg.OSS.End_point, config.Cfg.OSS.Access_key_id, config.Cfg.OSS.Access_key_secret)
	if err != nil {
		log.Println("创建实例Error:", err)
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Cfg.OSS.Bucket_name)
	if err != nil {
		log.Println("获取存储空间Error:", err)
		return
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile("files/"+filehash+fileSuffix, config.Cfg.App.Location+filename)
	if err != nil {
		log.Println("本地文件上传Error:", err)
		return
	}

}

//从oss下载文件
func DownLoadFromOSS(fileHash, fileType string) []byte {
	// 创建OSSClient实例。
	client, err := oss.New(config.Cfg.OSS.End_point, config.Cfg.OSS.Access_key_id, config.Cfg.OSS.Access_key_secret)
	if err != nil {
		log.Println("Error:", err)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Cfg.OSS.Bucket_name)
	if err != nil {
		log.Println("Error:", err)
	}

	// 下载文件到流。
	body, err := bucket.GetObject("files/" + fileHash + fileType)
	if err != nil {
		log.Println("Error:", err)
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("Error:", err)
	}

	return data
}

//从oss删除文件
func DeleteFromOSS(fileHash, fileType string) {
	// 创建OSSClient实例。
	client, err := oss.New(config.Cfg.OSS.End_point, config.Cfg.OSS.Access_key_id, config.Cfg.OSS.Access_key_secret)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Cfg.OSS.Bucket_name)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	err = bucket.DeleteObject("files/" + fileHash + fileType)
	if err != nil {
		log.Println("Error:", err)
		return
	}
}