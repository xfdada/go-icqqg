package config

import (
	"flag"
	"fmt"
	"gin-icqqg/utils/snow_id"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Server  *Server
	Redis   *Redis
	Mysql   *Mysql
	Jwt     *Jwt
	Logger  *Logger
	DB      *gorm.DB
	Upload  *Upload
	CaptCha *CaptCha
	AlySms  *AlySms
}

var AppConfig Config
var Snow *snow_id.Snow
var (
	port    string
	runMode string
	cfgpath string
)

func init() {
	err := setupFlag()
	if err != nil {
		fmt.Println(err)
	}
	err = initConfig()
	if err != nil {
		panic(err)
	}
	DB, err := NewDB()
	if err != nil {
		panic(err)
	}
	AppConfig.DB = DB
	Snow = snow_id.NewSnow(1)
}

func initConfig() error {
	err := NewConfig("config.yaml")
	if err != nil {
		return err
	}
	return nil
}
func NewConfig(config string) error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("./")
	if config != "" {
		vp.AddConfigPath(config)
	}
	vp.SetConfigType("yaml") // 设置配置文件类型格式为YAML
	err := vp.ReadInConfig()
	err = vp.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		if err = vp.Unmarshal(&AppConfig); err != nil {
			log.Printf("Config file changed filed: %s", err)
		}
	})
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&cfgpath, "config", "./", "配置文件的路径")
	flag.Parse()
	return nil
}
