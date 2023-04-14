package main

import (
	_ "cloud_storage/config"
	_ "cloud_storage/dao"
	"cloud_storage/router"
	"log"
)

func main() {
	//程序入口，一个项目 只有一个入口
	//web程序，http协议 ip port

	router.LoadRoute()
	router.Router.LoadHTMLGlob("template/*")
	router.Router.Static("/static", "./static")

	err := router.Router.Run()

	if err != nil {
		log.Fatal("服务器启动失败")
	}

}
