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
	DB       uint8
}

type Logger struct {
	AppPath    string
	SqlPath    string
	AppLevel   string
	SqlLevel   string
	ErrorPath  string
	ErrorLevel string
}
