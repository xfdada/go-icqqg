package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//Flow 流量来源分析
type Flow struct {
	Id        int64      `gorm:"column:id;type:int(11);" json:"id,omitempty"`
	PageNum   int64      `gorm:"column:page_num;type:int(11);" json:"page_num,omitempty"`
	IsNew     int64      `gorm:"column:is_new;type:tinyint(3);default:1;" json:"is_new,omitempty"`
	IP        string     `gorm:"column:ip;type:varchar(50);" json:"ip,omitempty"`
	GroupId   string     `gorm:"column:group_id;type:varchar(50);" json:"group_id,omitempty"`
	UserId    string     `gorm:"column:user_id;type:varchar(191);" json:"user_id,omitempty"`
	City      string     `gorm:"column:city;type:varchar(191);" json:"city,omitempty"`
	Origin    string     `gorm:"column:origin;type:varchar(2000);" json:"origin,omitempty"`
	Page      string     `gorm:"column:page;type:varchar(191);" json:"page,omitempty"`
	CreatedAt *LocalTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
}

//AddFlow 流量来源分析
type AddFlow struct {
	PageNum int64  ` json:"page_num,omitempty"`
	IsNew   int64  `json:"is_new,omitempty"`
	IP      string ` json:"ip,omitempty"`
	UserId  string ` json:"user_id,omitempty"`
	City    string `json:"city,omitempty"`
	Origin  string `json:"origin,omitempty"`
	Page    string ` json:"page,omitempty"`
	GroupId string `json:"group_id,omitempty"`
}

func NewFlow() *Flow {
	db.AutoMigrate(&Flow{})
	if !db.Migrator().HasTable(&Flow{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&Flow{})
	}
	return &Flow{}
}

func (f *Flow) TableName() string {
	return "table_flow"
}

func (f *Flow) List(c *gin.Context) {
	var List []Flow
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	Dbpage := Paginate(int(pageInt), int(pageSizeInt))
	err := Dbpage(db).Model(&Flow{}).Order("created_at desc").Find(&List).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		var Response map[string]interface{}
		var count int64
		Response = make(map[string]interface{})
		Response["code"] = 0
		Response["data"] = List
		db.Model(&Flow{}).Count(&count)
		Response["count"] = count
		Response["page"] = pageInt
		Response["pageSize"] = pageSizeInt
		c.JSON(200, Response)
	}
	c.Abort()
}

//UpOrAdd 新增或创建
func (f *Flow) UpOrAdd(add AddFlow) {
	now := time.Now()
	befor := now.Add(-2 * time.Hour)
	db.Model(&Flow{}).Where("user_id = ? and created_at >= ?", add.UserId, befor).First(f)
	fmt.Println(add)
	if f.UserId != "" {
		//更新
		f.PageNum += 1
		err := db.Model(&Flow{}).Where("user_id = ? ", add.UserId).Update("page_num", f.PageNum).Error
		config.ErrorLog(fmt.Sprintf("流量数据更新失败错误详情:%s", err))
	} else {
		var count int64
		db.Model(&Flow{}).Where("user_id = ? ", add.UserId).Count(&count)
		//创建
		f.Page = add.Page
		f.PageNum = 1
		f.GroupId = add.GroupId
		f.City = add.City
		f.UserId = add.UserId
		f.IP = add.IP
		f.Origin = add.Origin
		f.IsNew = 1 //新访客
		if count > 0 {
			f.IsNew = 2 //老访客
		}
		err := db.Model(&Flow{}).Create(&f).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("流量数据创建失败错误详情:%s", err))
		}
	}
	return
}

type Data struct {
	ID     int    `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	IP     string `json:"ip,omitempty"`
}

func (f *Flow) GetByHour(c *gin.Context) {
	var results []struct {
		Hour string `json:"hour,omitempty"`
		Uv   int    `json:"uv,omitempty"`
		Pv   int    `json:"pv,omitempty"`
		Ip   int    `json:"ip,omitempty"`
	}

	startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local) // 获取今天的起始时间
	db.Model(&Flow{}).Where("created_at >= ? AND created_at < ?", startTime, startTime.Add(24*time.Hour)).
		Select("DATE_FORMAT(created_at, '%H:00') AS hour,COUNT(ip) AS ip, COUNT(user_id) AS uv,SUM(page_num) AS pv").
		Group("hour").
		Order("hour").
		Scan(&results)
	c.JSON(200, results)
	c.Abort()
}
