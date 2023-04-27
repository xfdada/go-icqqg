package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

//ImMessage 聊天消息
type ImMessage struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"msg_id,omitempty"`
	MangerId  int64      `gorm:"column:manger_id;comment:管理员ID;type:varchar(191);" json:"manger_id,omitempty"`
	UserName  int64      `gorm:"column:username;comment:聊天用户;type:varchar(191);" json:"username,omitempty"`
	Avatar    string     `gorm:"avatar:username;comment:头像;type:varchar(191);" json:"avatar"`
	GroupId   string     `gorm:"column:group_id;comment:平台ID;type:varchar(191);" json:"groupid,omitempty"`
	GroupName string     `gorm:"column:group_name;comment:平台名称;type:varchar(191);" json:"groupname,omitempty"`
	UserId    string     `gorm:"column:user_id;comment:用户ID;type:varchar(191);" json:"id,omitempty"`
	Content   string     `gorm:"column:content;comment:消息内容;" json:"content,omitempty"`
	MsgType   string     `gorm:"column:msg_type;comment:消息内容;type:varchar(50);" json:"msgtype"`
	CreatedAt *LocalTime `gorm:"column:created_at;comment:创建时间;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;comment:更新时间;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;comment:删除时间;" json:"deleted_at,omitempty"`
}

type AddMessage struct {
	UserName int64  `gorm:"column:username;comment:管理员ID;" json:"username,omitempty"`
	MangerId int64  `gorm:"column:manger_id;comment:管理员ID;" json:"manger_id,omitempty"`
	GroupId  string `gorm:"column:group_id;comment:平台ID;" json:"group_id,omitempty"`
	UserId   string `gorm:"column:user_id;comment:用户ID;" json:"user_id,omitempty"`
	Message  string `gorm:"column:message;comment:消息内容;" json:"message,omitempty"`
}

func NewImMessage() *ImMessage {
	db.AutoMigrate(&ImMessage{})
	if !db.Migrator().HasTable(&ImMessage{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&ImMessage{})
	}
	return &ImMessage{}
}

func (msg *ImMessage) TableName() string {
	return "table_immessage"
}

//List 获取列表
func (msg *ImMessage) List(c *gin.Context) {
	var List []ImMessage
	pageInt := c.Query("page")
	pageNum := c.Query("pageNum")
	start := c.Query("start")
	end := c.Query("end")
	page, _ := strconv.ParseInt(pageInt, 10, 64)
	num, _ := strconv.ParseInt(pageNum, 10, 64)
	pages := Paginate(int(page), int(num))
	var total int64
	var err error
	if start != "" && end != "" {
		err = pages(db.Model(&ImMessage{})).Where("created_at BETWEEN ? AND ?", start, end).Find(&List).Error
		db.Model(&ImMessage{}).Where("created_at BETWEEN ? AND ?", start, end).Count(&total)
	} else if start != "" && end == "" {
		err = pages(db.Model(&ImMessage{})).Where("created_at > ? ", start).Find(&List).Error
		db.Model(&ImMessage{}).Where("created_at > ? ", start).Count(&total)
	} else if start == "" && end != "" {
		err = pages(db.Model(&ImMessage{})).Where("created_at < ? ", end).Find(&List).Error
		db.Model(&ImMessage{}).Where("created_at < ? ", end).Count(&total)
	} else {
		err = pages(db.Model(&ImMessage{})).Find(&List).Error
		db.Model(&ImMessage{}).Count(&total)
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

//Add 批量新增数据
func (msg *ImMessage) Add(add []ImMessage) {
	err := db.Model(&ImMessage{}).Create(add).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("前端用户新增失败错误详情是：%v", err))
	}
}

//GetUserList 管理员获取用聊天记录
func (msg *ImMessage) GetUserList(userId string, c *gin.Context) {
	var List []ImMessage
	err := db.Model(&ImMessage{}).Where("user_id = ?", userId).Find(&List).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "获取错误"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": List})
	}
	c.Abort()
}
