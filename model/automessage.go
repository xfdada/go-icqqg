package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
)

//AutoMessage 自动发送消息
type AutoMessage struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	IsDefault int64      `gorm:"column:is_default;type:tinyint(3);default:1;" json:"is_default,omitempty"`
	SendSort  int64      `gorm:"column:send_sort;type:int(11);" json:"send_sort,omitempty"`
	GroupId   string     `gorm:"column:group_id;type:varchar(255);" json:"group_id,omitempty"`
	Content   string     `gorm:"column:content;type:text;" json:"content,omitempty"`
	CreatedAt *LocalTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
}

type AddAutoMessage struct {
	IsDefault int64  `form:"is_default" json:"is_default,omitempty"` //是否默认
	SendSort  int64  `form:"send_sort" json:"send_sort,omitempty"`   //发送顺序
	GroupId   string `form:"group_id" json:"group_id,omitempty"`     //平台ID
	Content   string `form:"content" json:"content,omitempty"`       //发送内容
}

func NewAutoMessage() *AutoMessage {
	db.AutoMigrate(&AutoMessage{})
	if !db.Migrator().HasTable(&AutoMessage{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&AutoMessage{})
	}
	return &AutoMessage{}
}

func (msg *AutoMessage) TableName() string {
	return "table_automessage"
}

//List 获取自动发送消息列表
func (msg *AutoMessage) List(c *gin.Context) {
	var List []AutoMessage
	err := db.Model(&AutoMessage{}).Find(&List).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": List})
	}
	c.Abort()
}

//Add 新增消息
func (msg *AutoMessage) Add(add AddAutoMessage, c *gin.Context) {
	msg.IsDefault = add.IsDefault
	msg.GroupId = add.GroupId
	msg.SendSort = add.SendSort
	msg.Content = add.Content
	err := db.Model(&AutoMessage{}).Create(msg).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "新增成功"})
	}
	c.Abort()
}

//Edit 编辑消息
func (msg *AutoMessage) Edit(add AddAutoMessage, c *gin.Context) {
	id := c.Param("id")
	msg.IsDefault = add.IsDefault
	msg.GroupId = add.GroupId
	msg.SendSort = add.SendSort
	msg.Content = add.Content
	fmt.Println(add)
	err := db.Model(&AutoMessage{}).Where("id = ?", id).Updates(&msg).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "编辑失败"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "编辑成功"})
	}
	c.Abort()
}

//Delete 删除消息
func (msg *AutoMessage) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		err := db.Model(&AutoMessage{}).Where("id = ?", id).Delete(msg).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "删除失败"})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "删除成功"})
		}
	}
	c.Abort()
}

//Get 获取消息
func (msg *AutoMessage) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		err := db.Model(&AutoMessage{}).Where("id = ?", id).First(msg).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "删除失败"})
		} else {
			c.JSON(200, gin.H{"code": 200, "data": msg})
		}
	}
	c.Abort()
}
