package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
)

//Role 角色结构体
type Role struct {
	ID            int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	RoleName      string     `gorm:"column:role_name;type:varchar(50);" json:"role_name,omitempty"`            // 角色名称
	RoleTag       string     `gorm:"column:role_tag;type:varchar(50);" json:"role_tag,omitempty"`              // 角色标签
	MenuID        string     `gorm:"column:menu_id;type:varchar(255);" json:"menu_id,omitempty"`               // 菜单ID    用逗号分隔例如 1，2，3，4
	PermissionsID string     `gorm:"column:permissions_id;type:varchar(255);" json:"permissions_id,omitempty"` // 权限ID    用逗号分隔例如 1，2，3，4
	CreatedAt     *LocalTime `gorm:"column:created_at" json:"created_at,omitempty"`                            //创建时间
	UpdatedAt     *LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"`                            //更新时间
	DeletedAt     *LocalTime `gorm:"column:deleted_at" json:"-,omitempty"`                                     //删除时间
}

func NewRole() *Role {
	db.AutoMigrate(&Role{})
	if !db.Migrator().HasTable(&Role{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&Role{})
		db.Exec("ALTER TABLE admin_role ENGINE=InnoDB;")
	}
	return &Role{}
}
func (m *Role) TableName() string {
	return "admin_role"
}

//List 获取菜单列表展示的
//get
func (m *Role) List(c *gin.Context) {
	var RoleList []Role
	err := db.Model(&Role{}).Find(&RoleList).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": RoleList})
	}
	c.Abort()
}

//AddRole 新增角色
//params name,tag,menu_id,permissions_id
//post
func (m *Role) AddRole(params map[string]interface{}, c *gin.Context) {
	m.RoleName = params["name"].(string)
	m.RoleTag = params["tag"].(string)
	m.MenuID = params["menu_id"].(string)
	m.PermissionsID = params["permissions_id"].(string)
	err := db.Model(&Role{}).Create(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//DeleteRole 通过ID删除角色
//delete
func (m *Role) DeleteRole(id string, c *gin.Context) {
	err := db.Model(&Role{}).Where("id = ?", id).Delete(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//UpdateRole 修改角色
//params name,tag,menu_id,permissions_id
//put
func (m *Role) UpdateRole(id string, params map[string]interface{}, c *gin.Context) {
	m.RoleName = params["name"].(string)
	m.RoleTag = params["tag"].(string)
	m.MenuID = params["menu_id"].(string)
	m.PermissionsID = params["permissions_id"].(string)
	err := db.Model(&Role{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//GetOne 通过ID获取信息
//get
func (m *Role) GetOne(id string, c *gin.Context) {
	err := db.Model(&Role{}).Where("id = ?", id).First(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": m})
	}
	c.Abort()
}
