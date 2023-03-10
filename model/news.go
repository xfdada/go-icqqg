package model

import (
	"gin-icqqg/config/response"
	"github.com/gin-gonic/gin"
)

type News struct {
	Model
	Show       int    `json:"show,omitempty" gorm:"comment:是否显示 1是 2 否;type:int;size:2;" form:"show"`                                  //是否显示 1是 2 否
	Type       int    `json:"type,omitempty" gorm:"comment:产品类型;type:int;size:10;" form:"type"`                                        //产品类型
	SeeNum     int    `json:"see_num,omitempty"  gorm:"column:see_num;comment:查看数量;type:int;size:10;" form:"see_num"`                  // 查看数量
	Hot        int    `json:"hot,omitempty" gorm:"comment:是否热门 1是 2否;type:int;size:10;" form:"hot"`                                    // 是否热门 1是 2否
	Zan        int    `json:"zan,omitempty" gorm:"comment:点赞数量;type:int;size:10;" form:"zan"`                                          // 点赞数量
	CateId     int    `json:"cate_id,omitempty" gorm:"comment:归属类别;type:int;size:10;"  gorm:"column:cate_id" form:"cateId"`            //归属类别
	Tags       string `json:"tags,omitempty" gorm:"comment:新闻tag;type:string;size:255;" form:"tags"`                                   // 新闻tag
	Content    string `json:"content,omitempty" gorm:"comment:新闻内容" form:"content"`                                                    // 新闻内容
	SeoKeyword string `json:"seo_keyword,omitempty"  gorm:"column:seo_keyword;comment:新闻关键词;type:string;size:255;" form:"seo_keyword"` // 新闻关键词
	Title      string `json:"title,omitempty" gorm:"comment:新闻标题;type:string;size:255;" form:"title"`                                  // 新闻标题
	Thumb      string `json:"thumb,omitempty" gorm:"comment:新闻图片;type:string;size:255;" form:"thumb"`                                  // 新闻图片
	Desc       string `json:"desc,omitempty" gorm:"comment:新闻简介;type:string;size:255;" form:"desc"`                                    //新闻简介
}

func NewNews() *News {
	db.AutoMigrate(&News{})
	return &News{}
}

func (n News) TableName() string {
	return "table_news"
}

// GetById 通过ID获取内容
func (n *News) GetById(id string) (*News, int64) {
	res := db.Where("id=?", id).First(&n)
	return n, res.RowsAffected
}

func (n *News) AddNews() int64 {
	res := db.Create(n).RowsAffected
	return res
}

func (n *News) DeleteNews(c *gin.Context) {
	r := response.NewResponse(c)
	id := c.Param("id")
	res := db.Where("id = ?", id).Delete(&n)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.DeleteError)
	} else {
		r.SuccessResp("删除成功")
	}
	c.Abort()
}
