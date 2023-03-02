package config

import "time"

type Server struct {
	Port string
}

type Jwt struct {
	AppKey    string
	AppSecret string
	Expire    time.Duration
	Issuer    string
}

type Mysql struct {
	Username        string
	Password        string
	Host            string
	DBName          string
	Charset         string
	TablePrefix     string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

type Redis struct {
	Host     string
	Password string
	DB       int
}

type Logger struct {
	AppPath    string
	SqlPath    string
	AppLevel   string
	SqlLevel   string
	ErrorPath  string
	ErrorLevel string
}

type Upload struct {
	MaxSize int
	Path    string
	Url     string
	Ext     []string
}

type CaptCha struct {
	UseRedis   bool
	Height     int
	Width      int
	Length     int
	MaxSkew    float64
	DotCount   int
	PreKey     string
	Expiration time.Duration
}

// AlySms 阿里云短信配置
type AlySms struct {
	AliYunSmsAk string
	AliYunSmsAs string
	SingName    string
	Code        string
	Expiration  time.Duration
}
