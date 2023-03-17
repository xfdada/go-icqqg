package model

import (
	"database/sql/driver"
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var (
	db = config.AppConfig.DB
)

type LocalTime time.Time

type Model struct {
	ID        int            `gorm:"primary_key;autoincrement" json:"id"` //id
	CreatedAt int            `json:"-" gorm:"autoCreateTime"`             //创建时间
	UpdatedAt int            `json:"updated_at"`                          //更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type TableName struct {
	TableName string
}

type TableInfo struct {
	ColumnName    string `gorm:"column:column_name" json:"column_name"`
	ColumnDefault string `gorm:"column:column_default" json:"column_default"`
	IsNullable    string `gorm:"column:is_nullable = 'YES'" json:"is_nullable"`
	DataType      string `gorm:"column:data_type" json:"data_type"`
	ColumnType    string `gorm:"column:column_type" json:"column_type"`
	ColumnComment string `gorm:"column:column_comment" json:"column_comment"`
}

//Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.Unix() == zeroTime.Unix() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

//GetTable 获取数据库表名
func GetTable(c *gin.Context) {
	var table_list []TableName
	rows, err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = ?", "go_admin").Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
	}
	for rows.Next() {
		db.ScanRows(rows, &table_list)
	}
	c.JSON(200, gin.H{"table": table_list})
	c.Abort()
	return
}

//GetTableInfo 获取表详情

func MyTable(c *gin.Context) {
	var info []TableInfo
	row, _ := db.Raw("SELECT column_name, column_default, is_nullable = 'YES', data_type, column_type, column_comment FROM information_schema.columns WHERE table_schema = ? AND table_name = ?", "go_admin", "table_news").Rows()
	for row.Next() {
		db.ScanRows(row, &info)
	}
	//row.Scan(&info.ColumnName, &info.ColumnDefault, &info.IsNullable, &info.DataType, &info.ColumnType, &info.ColumnComment)
	c.JSON(200, gin.H{"data": info})
	c.Abort()
	return
}
