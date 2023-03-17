package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
)

//Permissions 权限表
type Permissions struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	Name       string     `gorm:"column:name;type:varchar(50);" json:"name,omitempty"`               //权限英文名
	Tag        string     `gorm:"column:tag;type:varchar(50);" json:"tag,omitempty"`                 //权限名称
	HttpMethod string     `gorm:"column:http_method;type:varchar(50);" json:"http_method,omitempty"` //请求方法  空为任何方法 ，GET，POST，PUT，DELETE，PATCH，OPTIONS
	HttpPath   string     `gorm:"column:http_path;type:varchar(255);" json:"http_path,omitempty"`    //请求路径
	ParentId   int64      `gorm:"column:parent_id;type:varchar(50);" json:"parent_id,omitempty"`     //父级ID
	CreatedAt  *LocalTime `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  *LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  *LocalTime `gorm:"column:deleted_at" json:"-,omitempty"`
}
type PermissionTree struct {
	Data *Permissions     `json:"data"`
	Sub  []PermissionTree `json:"sub,omitempty"`
}

func NewPermissions() *Permissions {
	db.AutoMigrate(&Permissions{})
	if !db.Migrator().HasTable(&Permissions{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&Permissions{})
	}
	return &Permissions{}
}
func (m *Permissions) TableName() string {
	return "admin_permissions"
}

//List 获取权限列表展示的
//get
func (m *Permissions) List(c *gin.Context) {
	var PermissionsList []Permissions
	row, err := db.Model(&Permissions{}).Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		for row.Next() {
			db.ScanRows(row, &PermissionsList)
		}
		c.JSON(200, gin.H{"code": 200, "data": PermissionsList})
	}
	c.Abort()
}

//AddPermissions 新增权限
//params parent_id,sort,title,url,show
//post
func (m *Permissions) AddPermissions(params map[string]interface{}, c *gin.Context) {
	m.Name = params["name"].(string)
	m.Tag = params["tag"].(string)
	m.HttpPath = params["http_path"].(string)
	m.HttpMethod = params["http_method"].(string)
	m.ParentId = params["parent_id"].(int64)
	err := db.Model(&Permissions{}).Create(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//DeleteMenu 通过ID删除权限
//delete
func (m *Permissions) DeleteMenu(id string, c *gin.Context) {
	err := db.Model(&Permissions{}).Where("id = ?", id).Delete(&m).Error
	err = db.Model(&Permissions{}).Where("parent_id = ?", id).Delete(&m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//UpdateMenu 修改权限
//params parent_id,sort,title,url,show
//put
func (m *Permissions) UpdateMenu(id string, params map[string]interface{}, c *gin.Context) {
	m.Name = params["name"].(string)
	m.Tag = params["tag"].(string)
	m.HttpPath = params["http_path"].(string)
	m.HttpMethod = params["http_method"].(string)
	m.ParentId = params["parent_id"].(int64)
	err := db.Model(&Permissions{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//GetOne 通过ID获取权限信息
//get
func (m *Permissions) GetOne(id string, c *gin.Context) {
	err := db.Model(&Permissions{}).Where("id = ?", id).First(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": m})
	}
	c.Abort()
}

//GetParent 获取顶级权限信息列表
//get
func (m *Permissions) GetParent(c *gin.Context) {
	var PermissionsList []Permissions
	rows, err := db.Model(&Permissions{}).Where("parent_id = 0").Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		for rows.Next() {
			db.ScanRows(rows, &PermissionsList)
		}
		c.JSON(200, gin.H{"code": 200, "data": PermissionsList})
	}
	c.Abort()
}

//GetPermissionTree 获取权限层级信息列表
//get
func (m *Permissions) GetPermissionTree(c *gin.Context) {
	var Permission []PermissionTree
	var PermissionParent []*Permissions
	rows, err := db.Model(&Permissions{}).Where("parent_id = 0").Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		for rows.Next() {
			db.ScanRows(rows, &PermissionParent)
		}
		Permission = m.createdTree(PermissionParent)
		c.JSON(200, gin.H{"code": 200, "data": &Permission})
	}
	c.Abort()
}

//createdTree 创建无限极树
func (m *Permissions) createdTree(parent []*Permissions) []PermissionTree {
	var List []PermissionTree
	for _, v := range parent {
		var count int64
		db.Model(&Permissions{}).Where("parent_id = ?", v.ID).Count(&count)
		node := PermissionTree{Data: v}
		if count > 0 {
			var Sub []*Permissions
			db.Preload("Permissions").Where("parent_id = ?", v.ID).Find(&Sub)
			node.Sub = m.createdTree(Sub)
		}
		List = append(List, node)
	}
	return List
}

//fullDel 获取完全删除对象的id切片
func (m *Permissions) fullDel(parentId int64) []int64 {
	var Ids []int64
	db.Model(&Permissions{}).Pluck("id", &Ids)
	stack := make([]int64, 0)
	stack = append(stack, parentId)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		Ids = append(Ids, node)

		var children []int64
		db.Model(&Permissions{}).Where("parent_id = ?", node).Pluck("id", &children)
		for _, child := range children {
			stack = append(stack, child)
		}
	}

	return Ids
}
