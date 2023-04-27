package model

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

type News struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement;type:int(11);" json:"id,omitempty"`
	Show       int        `json:"show,omitempty" gorm:"comment:是否显示 1是 2 否;type:tinyint(3);" form:"show"`                               //是否显示 1是 2 否
	Type       int        `json:"type,omitempty" gorm:"comment:产品类型;type:int(3);" form:"type"`                                          //产品类型
	SeeNum     int        `json:"see_num,omitempty"  gorm:"column:see_num;comment:查看数量;type:int(11);" form:"see_num"`                   // 查看数量
	Hot        int        `json:"hot,omitempty" gorm:"comment:是否热门 1是 2否;type:int(11);" form:"hot"`                                     // 是否热门 1是 2否
	Zan        int        `json:"zan,omitempty" gorm:"comment:点赞数量;type:int(11);" form:"zan"`                                           // 点赞数量
	CateId     int        `json:"cate_id,omitempty" gorm:"comment:归属类别;type:int(11);"  gorm:"column:cate_id" form:"cate_id"`            //归属类别
	Tags       string     `json:"tags,omitempty" gorm:"comment:新闻tag;type:varchar(255);" form:"tags"`                                   // 新闻tag
	Content    string     `json:"content,omitempty" gorm:"comment:新闻内容" form:"content"`                                                 // 新闻内容
	SeoKeyword string     `json:"seo_keyword,omitempty"  gorm:"column:seo_keyword;comment:新闻关键词;type:varchar(255);" form:"seo_keyword"` // 新闻关键词
	Title      string     `json:"title,omitempty" gorm:"comment:新闻标题;type:varchar(255);" form:"title"`                                  // 新闻标题
	Thumb      string     `json:"thumb,omitempty" gorm:"comment:新闻图片;type:varchar(255);" form:"thumb"`                                  // 新闻图片
	Desc       string     `json:"desc,omitempty" gorm:"comment:新闻简介;type:varchar(255);" form:"description"`                             //新闻简介
	CreatedAt  *LocalTime `gorm:"column:created_at;autoCreateTime;" json:"created_at,omitempty"`
	UpdatedAt  *LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  *LocalTime `gorm:"column:deleted_at" json:"-,omitempty"`
}
type AddNews struct {
	Show       int    `json:"show,omitempty"  form:"show"`                //是否显示 1是 2 否
	Type       int    `json:"type,omitempty"  form:"type"`                //产品类型
	SeeNum     int    `json:"see_num,omitempty"   form:"see_num"`         // 查看数量
	Hot        int    `json:"hot,omitempty"  form:"hot"`                  // 是否热门 1是 2否
	Zan        int    `json:"zan,omitempty" form:"zan"`                   // 点赞数量
	CateId     int    `json:"cate_id,omitempty"  form:"cate_id"`          //归属类别
	Tags       string `json:"tags,omitempty"  form:"tags"`                // 新闻tag
	Content    string `json:"content,omitempty"  form:"content"`          // 新闻内容
	SeoKeyword string `json:"seo_keyword,omitempty"   form:"seo_keyword"` // 新闻关键词
	Title      string `json:"title,omitempty"  form:"title"`              // 新闻标题
	Thumb      string `json:"thumb,omitempty"  form:"thumb"`              // 新闻图片
	Desc       string `json:"description,omitempty" form:"description"`
}

func NewNews() *News {
	db.AutoMigrate(&News{})
	if !db.Migrator().HasTable(&News{}) {
		db.Set("gorm:ENGINE", "InnoDB").Migrator().CreateTable(&News{})
	}
	return &News{}
}

func (n News) TableName() string {
	return "table_news"
}

//List 获取产品列表展示的
//get
func (m *News) List(c *gin.Context) {
	var NewsList []News
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	dbpage := Paginate(int(pageInt), int(pageSizeInt))
	err := dbpage(db).Model(&News{}).Find(&NewsList).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		var count int64
		var Res map[string]interface{}
		Res = make(map[string]interface{})
		Res["data"] = NewsList
		Res["page"] = page
		db.Model(&News{}).Count(&count)
		Res["count"] = count
		Res["code"] = 0
		Res["msg"] = ""
		c.JSON(200, Res)
	}
	c.Abort()
}

//AddNews 新增新闻
//post
func (m *News) AddNews(c *gin.Context) {
	var add AddNews
	err := c.ShouldBind(&add)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": fmt.Sprintf("%v", err)})
	} else {
		m.Show = add.Show
		m.Type = add.Type
		m.SeeNum = add.SeeNum
		m.Hot = add.Hot
		m.Zan = add.Zan
		m.CateId = add.CateId
		m.Tags = add.Tags
		m.Content = add.Content
		m.SeoKeyword = add.SeoKeyword
		m.Title = add.Title
		m.Thumb = add.Thumb
		m.Desc = add.Desc
		err = db.Model(&News{}).Create(&m).Error
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "failed"})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "success"})
		}
	}
	c.Abort()
}

//DeleteProduct 通过ID删除产品
//delete
func (m *News) DeleteNews(id string, c *gin.Context) {
	err := db.Model(&News{}).Where("id = ?", id).Delete(m).Error
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
func (m *News) UpdateNews(id string, c *gin.Context) {
	var add AddNews
	err := c.ShouldBind(&add)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": fmt.Sprintf("%v", err)})
	} else {
		m.Show = add.Show
		m.Type = add.Type
		m.SeeNum = add.SeeNum
		m.Hot = add.Hot
		m.Zan = add.Zan
		m.CateId = add.CateId
		m.Tags = add.Tags
		m.Content = add.Content
		m.SeoKeyword = add.SeoKeyword
		m.Title = add.Title
		m.Thumb = add.Thumb
		m.Desc = add.Desc
		if err != nil {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(200, gin.H{"code": 500, "msg": "failed"})
		} else {
			err := db.Model(&News{}).Where("id = ?", id).Updates(m).Error
			if err != nil {
				config.ErrorLog(fmt.Sprintf("%v", err))
				c.JSON(200, gin.H{"code": 500, "msg": "failed"})
			} else {
				c.JSON(200, gin.H{"code": 200, "msg": "success"})
			}
		}
	}
	c.Abort()
}

//GetOne 通过ID获取信息
//get
func (m *News) GetOne(id string, c *gin.Context) {
	err := db.Model(&News{}).Where("id = ?", id).First(m).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(200, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": m})
	}
	c.Abort()
}
