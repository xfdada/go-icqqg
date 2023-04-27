package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

//ImOffer 离线留言
type ImOffer struct {
	Id        int64      `gorm:"column:id;type:int(11);" json:"id,omitempty"`
	GroupId   string     `gorm:"column:group_id;type:varchar(90);" json:"group_id,omitempty"`
	Name      string     `gorm:"column:name;type:varchar(191);" json:"name,omitempty"`
	Email     string     `gorm:"column:email;type:varchar(191);" json:"email,omitempty"`
	Content   string     `gorm:"column:content;type:text;" json:"content,omitempty"`
	CreatedAt *LocalTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
}
type AddImOffer struct {
	GroupId string `form:"group_id" json:"group_id,omitempty"`
	Name    string `form:"name" json:"name,omitempty"`
	Email   string `form:"email" json:"email,omitempty"`
	Content string `form:"content" json:"content,omitempty"`
}

func NewImOffer() *ImOffer {
	db.AutoMigrate(&ImOffer{})
	if !db.Migrator().HasTable(&ImOffer{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&ImOffer{})
	}
	return &ImOffer{}
}

func (im *ImOffer) TableName() string {
	return "table_imoffer"
}

//List 获取离线留言列表
//参数 page   pageSize  group_id
func (im *ImOffer) List(c *gin.Context) {
	var List []ImOffer
	var count int64
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	groupId := c.Query("group_id")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	dbPage := Paginate(int(pageInt), int(pageSizeInt))
	var err error
	if groupId != "" {
		err = dbPage(db).Model(&ImOffer{}).Where("group_id = ?", groupId).Order("created_at desc").Find(&List).Error
		db.Model(&ImOffer{}).Where("group_id = ?", groupId).Count(&count)
	} else {
		err = dbPage(db).Model(&ImOffer{}).Order("created_at desc").Find(&List).Error
		db.Model(&ImOffer{}).Count(&count)
	}
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		var Response map[string]interface{}
		Response = make(map[string]interface{})
		Response["code"] = 0
		Response["data"] = List
		Response["count"] = count
		c.JSON(200, Response)
	}
	c.Abort()
}

//Add 新增消息
func (im *ImOffer) Add(c *gin.Context) {
	var add AddImOffer
	err := c.ShouldBind(&add)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "参数有误"})
		c.Abort()
		return
	}
	im.GroupId = add.GroupId
	im.Name = add.Name
	im.Email = add.Email
	im.Content = add.Content
	err = db.Model(&ImOffer{}).Create(im).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "新增成功"})
	}
	c.Abort()
}

//Edit 编辑消息
//func (msg *ImOffer) Edit(c *gin.Context) {
//	id := c.Param("id")
//	var add AddImOffer
//	err:= c.ShouldBind(&add)
//	if err != nil{
//		c.JSON(200, gin.H{"code": 500, "msg": "新增失败"})
//		c.Abort()
//		return
//	}
//	msg.GroupId = add.GroupId
//	msg.Name = add.Name
//	msg.Email = add.Email
//	msg.Content = add.Content
//	err = db.Model(&ImOffer{}).Where("id = ?", id).Updates(&msg).Error
//	if err != nil {
//		config.ErrorLog(fmt.Sprintf("%v", err))
//		c.JSON(200, gin.H{"code": 500, "msg": "编辑失败"})
//	} else {
//		c.JSON(200, gin.H{"code": 200, "msg": "编辑成功"})
//	}
//	c.Abort()
//}
