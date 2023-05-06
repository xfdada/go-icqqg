package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ImCode struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	Group     string     `gorm:"column:group;comment:平台;type:varchar(191);" json:"group,omitempty"`        //平台
	GroupId   string     `gorm:"column:group_id;comment:平台ID;type:varchar(191);" json:"groupid,omitempty"` //平台ID
	Manger    string     `gorm:"column:manger;comment:管理员;type:varchar(191);" json:"manger,omitempty"`     //管理员
	CreatedAt *LocalTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	DeletedAt *LocalTime `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
}
type AddImCode struct {
	Group   string `json:"group" form:"group"`     //平台
	GroupId string `json:"groupid" form:"groupid"` //平台ID
	Manger  string `json:"manger" form:"manger"`   //管理员
}

func NewImCode() *ImCode {
	db.AutoMigrate(&ImCode{})
	if !db.Migrator().HasTable(&ImCode{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&ImCode{})
	}
	return &ImCode{}
}

func (msg *ImCode) TableName() string {
	return "table_imcode"
}

//List 获取客服代码列表
//get
func (m *ImCode) List(manageId string, c *gin.Context) {
	var NewsList []ImCode
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	dbpage := Paginate(int(pageInt), int(pageSizeInt))
	var err error
	if manageId == "10086" {
		err = dbpage(db).Model(&ImCode{}).Find(&NewsList).Error
	} else {
		err = dbpage(db).Model(&ImCode{}).Where("manger = ?", manageId).Find(&NewsList).Error
	}
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		var count int64
		var Res map[string]interface{}
		Res = make(map[string]interface{})
		Res["data"] = NewsList
		Res["page"] = page
		if manageId == "10086" {
			db.Model(&ImCode{}).Count(&count)
		} else {
			db.Model(&ImCode{}).Where("manger = ?", manageId).Count(&count)
		}
		Res["count"] = count
		Res["code"] = 0
		Res["msg"] = ""
		c.JSON(200, Res)
	}
	c.Abort()
}

//AddImCode 新增代码
//post
func (m *ImCode) AddImCode(ManageId string, c *gin.Context) {
	var add AddImCode
	err := c.ShouldBind(&add)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": fmt.Sprintf("%v", err)})
	} else {
		m.Group = add.Group
		m.GroupId = add.GroupId
		m.Manger = ManageId
		err = db.Model(&ImCode{}).Create(&m).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "failed"})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "success"})
		}
	}
	c.Abort()
}

func (m *ImCode) Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		err := db.Model(&ImCode{}).Where("id = ?", id).Delete(m).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "failed"})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "success"})
		}
	}
	c.Abort()
}

//EditImCode 更新代码
//PUT
func (m *ImCode) EditImCode(c *gin.Context) {
	err := c.ShouldBind(&m)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": fmt.Sprintf("%v", err)})
	} else {
		err = db.Model(&ImCode{}).Updates(&m).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "failed"})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "success"})
		}
	}
	c.Abort()
}
