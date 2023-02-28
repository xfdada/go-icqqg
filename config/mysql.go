package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	SqlLog(s)
}

func NewDB() (*gorm.DB, error) {
	newLogger := logger.New(
		Writer{}, // io writer
		logger.Config{
			SlowThreshold: 100 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // Log level
			Colorful:      false,                  // 禁用彩色打印
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=Local",
		AppConfig.Mysql.Username, AppConfig.Mysql.Password, AppConfig.Mysql.Host, AppConfig.Mysql.DBName, AppConfig.Mysql.Charset, AppConfig.Mysql.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		msg := "初始化连接数据库失败，错误详情是err:" + fmt.Sprintf("%v\n", err)
		fmt.Println(msg)
		return nil, err
	}
	sqlDB, err1 := db.DB()
	if err1 != nil {
		msg := "初始化连接数据库失败，错误详情是err:" + fmt.Sprintf("%v\n", err1)
		fmt.Println(msg)
		return nil, err1
	}
	sqlDB.SetMaxIdleConns(AppConfig.Mysql.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(AppConfig.Mysql.ConnMaxIdleTime * time.Second) //连接池空闲的最长时间
	sqlDB.SetConnMaxLifetime(AppConfig.Mysql.ConnMaxLifetime * time.Second) //可重用连接的最长时间
	sqlDB.SetMaxOpenConns(AppConfig.Mysql.MaxOpenConns)
	return db, nil
}
