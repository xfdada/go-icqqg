package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
)

//Product 产品表
type Product struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	Name       string     `gorm:"column:name;type:varchar(50);" json:"name,omitempty"` //
	Title       string     `gorm:"column:title;type:varchar(50);" json:"title,omitempty"` //
	Version       string     `gorm:"column:version;type:varchar(50);" json:"version,omitempty"` //
	CreatedAt  *LocalTime `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  *LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  *LocalTime `gorm:"column:deleted_at" json:"-,omitempty"`
}

type AddProduct struct {
    Name       string     `form:"name" json:"name,omitempty"`//名称
    Title       string     `form:"title" json:"title,omitempty"`//标题
    Version       string     `form:"version" json:"version,omitempty"`//型号

}
func NewProduct() *Product {
	db.AutoMigrate(&Product{})
	if !db.Migrator().HasTable(&Product{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&Product{})
	}
	return &Product{}
}
func (m *Product) TableName() string {
	return "table_product"
}

//List 获取产品列表展示的
//get
func (m *Product) List(c *gin.Context) {
	var ProductList []Product
	err := db.Model(&Product{}).Find(&ProductList).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": ProductList})
	}
	c.Abort()
}

//AddProduct 新增产品
//post
func (m *Product) AddProduct( c *gin.Context) {
	var add AddProduct
	err := c.ShouldBind(&add)
	 m.Name = add.Name
	 m.Title = add.Title
	 m.Version = add.Version
	err = db.Model(&Product{}).Create(&m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}


//DeleteProduct 通过ID删除产品
//delete
func (m *Product) DeleteProduct(id string, c *gin.Context) {
	err := db.Model(&Product{}).Where("id = ?", id).Delete(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
	c.Abort()
}

//UpdateProduct 修改产品
//put
func (m *Product) UpdateProduct(id string,c *gin.Context) {
    var add AddProduct
	err := c.ShouldBind(&add)
	 m.Name = add.Name
	 m.Title = add.Title
	 m.Version = add.Version
    if err != nil {
    		config.ErrorLog(fmt.Sprintf("%v", err))
    		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
    }else{
        err := db.Model(&Product{}).Where("id = ?", id).Updates(m).Error
        if err != nil {
            config.ErrorLog(fmt.Sprintf("%v", err))
            c.JSON(200, gin.H{"code": 500, "msg": "failed"})
        } else {
            c.JSON(200, gin.H{"code": 200, "msg": "success"})
        }
    }

	c.Abort()
}

//GetOne 通过ID获取信息
//get
func (m *Product) GetOne(id string, c *gin.Context) {
	err := db.Model(&Product{}).Where("id = ?", id).First(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": m})
	}
	c.Abort()
}
