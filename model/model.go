package model

import (
	"gin-icqqg/config"
	"gorm.io/gorm"
	"time"
)

var (
	db = config.AppConfig.DB
)

type LocalTime time.Time

type Model struct {
	ID        int            `gorm:"primary_key;autoincrement" json:"id"` //id
	CreatedAt *LocalTime     `json:"created_at" gorm:"autoCreateTime"`    //创建时间
	UpdatedAt *LocalTime     `json:"updated_at"`                          //更新时间
	DeletedAt gorm.DeletedAt `json:"-"gorm:"index"`
}
