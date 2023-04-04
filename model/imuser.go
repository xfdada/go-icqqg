package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

//前端聊天客户

type ImUser struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	Manage    string     `gorm:"column:manage;comment:归属管理员;type:varchar(50);" json:"manage,omitempty"`      //归属管理员
	Group     string     `gorm:"column:group;comment:归属站点;type:varchar(50);" json:"group,omitempty"`         //归属站点
	GroupId   string     `gorm:"column:group_id;comment:归属站点;type:varchar(50);" json:"group_id,omitempty"`   //归属站点
	UserId    string     `gorm:"column:user_id;comment:用户ID;type:varchar(191);" json:"user_id,omitempty"`    //用户ID
	UserName  string     `gorm:"column:user_name;comment:用户名;type:varchar(191);" json:"user_name,omitempty"` //用户名
	Avatar    string     `gorm:"column:avatar;comment:头像;type:varchar(191);" json:"avatar,omitempty"`        //头像
	IP        string     `gorm:"column:ip;comment:IP地址;type:varchar(50);" json:"ip,omitempty"`               //IP地址
	CreatedAt *LocalTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
}

type AddImUser struct {
	Manage   string `form:"manage" json:"manage,omitempty"` //归属管理员
	Group    string `form:"group" json:"group,omitempty"`   //归属站点
	GroupId  string `form:"group_id" json:"group_id,omitempty"`
	UserId   string `form:"user_id" json:"user_id,omitempty"`     //用户ID
	UserName string `form:"user_name" json:"user_name,omitempty"` //用户名
	Avatar   string `form:"avatar" json:"avatar,omitempty"`       //头像
	IP       string `form:"ip" json:"ip,omitempty"`               //IP地址
}

func NewImUser() *ImUser {
	db.AutoMigrate(&ImUser{})
	if !db.Migrator().HasTable(&ImUser{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&ImUser{})
	}
	return &ImUser{}
}

func (im *ImUser) TableName() string {
	return "table_im"
}

//List 后台管理超级管理人员进行查看
func (im *ImUser) List(c *gin.Context) {
	var List []ImUser
	pageInt := c.Query("page")
	pageNum := c.Query("pageNum")
	start := c.Query("start")
	end := c.Query("end")
	manger := c.Query("manger")
	var query *gorm.DB
	page, _ := strconv.ParseInt(pageInt, 10, 64)
	num, _ := strconv.ParseInt(pageNum, 10, 64)
	pages := Paginate(int(page), int(num))
	var total int64
	var err error
	query = pages(db.Model(&ImUser{}))
	if manger != "" {
		query.Where("manger = ? ", manger)
	}
	if start != "" && end != "" {
		err = query.Where("created_at BETWEEN ? AND ?", start, end).Find(&List).Error
		query.Where("created_at BETWEEN ? AND ?", start, end).Count(&total)
	} else if start != "" && end == "" {
		err = query.Where("created_at > ? ", start).Find(&List).Error
		query.Where("created_at > ? ", start).Count(&total)
	} else if start == "" && end != "" {
		err = query.Where("created_at < ? ", end).Find(&List).Error
		query.Where("created_at < ? ", end).Count(&total)
	} else {
		err = query.Find(&List).Error
		query.Count(&total)
	}

	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "获取错误"})
	} else {
		PageList := map[string]interface{}{}
		PageList["data"] = List
		PageList["total"] = total
		PageList["page"] = page
		PageList["pageNum"] = num
		c.JSON(200, gin.H{"code": 200, "data": PageList})
	}
	c.Abort()
}

//Add 新增用户
func (im *ImUser) Add(user AddImUser) {
	im.IP = user.IP
	im.Manage = user.Manage
	im.Group = user.Group
	im.UserId = user.UserId
	im.GroupId = user.GroupId
	err := db.Model(&ImUser{}).Create(&im).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("前端用户新增失败错误详情是：%v", err))
	}
}

type Result struct {
	Group   string
	GroupId string
}

// GetFriendList 获取好友列表
//通过管理员账号获取 对应好友列表
func (im *ImUser) GetFriendList(manger string, c *gin.Context) {
	var result []Result
	err := db.Model(&ImUser{}).Group("group").Select("group", "group_id").Scan(&result).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("管理员获取好友列表错误：%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "获取错误"})
	} else {
		var List []map[string]interface{}
		for k, v := range result {
			var Friend []ImUser
			db.Model(&ImUser{}).Where("group = ?", v.Group).Find(&Friend)
			List[k] = map[string]interface{}{"groupname": v.Group, "id": v.GroupId, "list": Friend}
		}
		c.JSON(200, gin.H{"code": 200, "data": List})
	}
	c.Abort()
}
