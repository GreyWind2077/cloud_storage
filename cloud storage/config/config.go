package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

type config struct {
	Title  string
	DB     mysql `toml:"mysql"`
	App    app
	Redis  redis
	Server server
	QQ     qq
	OSS    oss
}

type app struct {
	Location  string
	Page_Size int
}

type server struct {
	Port          int
	Read_timeout  time.Duration
	Write_timeout time.Duration
}

type mysql struct {
	User     string
	Password string
	Host     string
	Name     string
}

type redis struct {
	Host  string
	Index string
}

type qq struct {
	App_id       string
	Redirect_url string
	App_key      string
}

type oss struct {
	Access_key_id     string
	Access_key_secret string
	End_point         string
	Bucket_name       string
}

var Cfg *config

func init() {
	Cfg = new(config)

	Cfg.App.Location, _ = os.Getwd()

	_, err := toml.DecodeFile("config/config.toml", &Cfg)

	if err != nil {
		log.Fatal("读取配置文件出错")
		panic(err)
	}
	log.Printf("全局信息: %+v\n", Cfg.Title)

	log.Printf("App信息：%+v\n", Cfg.App)

	log.Printf("Mysql配置：%+v\n", Cfg.DB)

	log.Printf("Server信息：%+v\n", Cfg.Server)

	log.Printf("Redis主从：%+v\n", Cfg.Redis)

	log.Printf("OSS信息：%+v\n", Cfg.OSS)

	log.Printf("QQ信息：%+v\n", Cfg.QQ)

}
