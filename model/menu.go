package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
)

//Menu 后台菜单
type Menu struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id"`
	ParentID  int64      `gorm:"column:parent_id;type:int(11);" json:"parent_id"` // 父级ID
	Sort      int64      `gorm:"column:sort;type:int(11);" json:"sort"`           //  菜单排序
	Title     string     `gorm:"column:title;type:varchar(50);" json:"title"`     // 菜单名称
	Url       string     `gorm:"column:url;type:varchar(50);" json:"url"`         // 路由地址
	Show      int64      `gorm:"column:show;type:tinyint(3);" json:"show"`        // 是否展示 1是 2否
	CreatedAt *LocalTime `json:"created_at"`
	UpdatedAt *LocalTime `json:"updated_at"`
	DeletedAt *LocalTime `json:"-"`
}
type MenuTree struct {
	Data *Menu      `json:"data"`
	Sub  []MenuTree `json:"sub,omitempty"`
}

func NewMenu() *Menu {
	db.AutoMigrate(&Menu{})
	if !db.Migrator().HasTable(&Menu{}) {
		db.Set("gorm:admin_menu", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
	}
	return &Menu{}
}

func (m *Menu) TableName() string {
	return "admin_menu"
}

//List 获取菜单列表展示的
//get
func (m *Menu) List(c *gin.Context) {
	var MenuList []Menu
	row, err := db.Model(&Menu{}).Where("`show` = 1").Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		for row.Next() {
			db.ScanRows(row, &MenuList)
		}
		c.JSON(200, gin.H{"code": 200, "data": MenuList})
	}
	c.Abort()
}

//AddMenu 新增菜单
//params parent_id,sort,title,url,show
//post
func (m *Menu) AddMenu(params map[string]interface{}, c *gin.Context) {
	m.Show = params["show"].(int64)
	m.ParentID = params["parent_id"].(int64)
	m.Sort = params["sort"].(int64)
	m.Title = params["title"].(string)
	m.Url = params["url"].(string)
	err := db.Model(&Menu{}).Create(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//DeleteMenu 通过ID删除菜单 如果是顶级菜单，那么子菜单也会被删除
//delete
func (m *Menu) DeleteMenu(id string, c *gin.Context) {
	err := db.Model(&Menu{}).Where("id = ?", id).Delete(&m).Error
	err = db.Model(&Menu{}).Where("parent_id = ?", id).Delete(&m).Error // 删除子菜单
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//UpdateMenu 修改菜单
//params parent_id,sort,title,url,show
//put
func (m *Menu) UpdateMenu(id string, params map[string]interface{}, c *gin.Context) {
	m.Show = params["show"].(int64)
	m.ParentID = params["parent_id"].(int64)
	m.Sort = params["sort"].(int64)
	m.Title = params["title"].(string)
	m.Url = params["url"].(string)
	err := db.Model(&Menu{}).Where("id = ?", id).Updates(m).Error
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
func (m *Menu) GetOne(id string, c *gin.Context) {
	err := db.Model(&Menu{}).Where("id = ?", id).First(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": m})
	}
	c.Abort()
}

//GetParent 获取顶级菜单信息列表
//get
func (m *Menu) GetParent(c *gin.Context) {
	var MenuList []Menu
	rows, err := db.Model(&Menu{}).Where("parent_id = 0").Rows()
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		for rows.Next() {
			db.ScanRows(rows, &MenuList)
		}
		c.JSON(200, gin.H{"code": 200, "data": MenuList})
	}
	c.Abort()
}

//GetMenuTree 获取菜单层级信息列表
//get
func (m *Menu) GetMenuTree(c *gin.Context) {
	var Menus []MenuTree
	var MenuParent []*Menu
	err := db.Model(&Menu{}).Where("parent_id = 0").Find(&MenuParent).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		Menus = m.createdTree(MenuParent)
		c.JSON(200, gin.H{"code": 200, "data": &Menus})
	}
	c.Abort()
}

//createdTree 创建无限极树
func (m *Menu) createdTree(parent []*Menu) []MenuTree {
	var List []MenuTree
	for _, v := range parent {
		var count int64
		db.Model(&Menu{}).Where("parent_id = ?", v.ID).Count(&count)
		node := MenuTree{Data: v}
		if count > 0 {
			var Sub []*Menu
			db.Preload("Menu").Where("parent_id = ?", v.ID).Find(&Sub)
			node.Sub = m.createdTree(Sub)
		}
		List = append(List, node)
	}
	return List
}
