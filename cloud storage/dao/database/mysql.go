package database

import (
	"cloud_storage/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func init() {

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.Host,
		config.Cfg.DB.Name,
	)
	DB, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal("数据库初始化错误", err)
	}
	DB.DB().SetMaxIdleConns(5)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Minute * 3)
	DB.DB().SetConnMaxIdleTime(time.Minute * 1)

	err = DB.DB().Ping()
	if err != nil {
		log.Println("数据库无法连接", err)
		_ = DB.Close()
		panic(err)
	} else {
		log.Println("数据库连接成功")
	}

}
