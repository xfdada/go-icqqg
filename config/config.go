package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Server *Server
	Redis  *Redis
	Mysql  *Mysql
	Jwt    *Jwt
	Logger *Logger
	DB     *gorm.DB
}

var AppConfig Config

func init() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
	DB, err := NewDB()
	if err != nil {
		panic(err)
	}
	AppConfig.DB = DB
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
		log.Printf("Config file changed: %s", e.Name)
		if err = vp.Unmarshal(&AppConfig); err != nil {
			log.Printf("Config file changed filed: %s", err)
		}
	})
	return nil
}
