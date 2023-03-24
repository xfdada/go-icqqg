package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type News struct {
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
	Desc       string `json:"desc,omitempty" gorm:"comment:新闻简介;type:string;size:255;" form:"desc"`
}

func NewNews() *News {
	return &News{}
}

//List 数据列表
//@Tags 后台产品模块
//@Summary 产品列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.Product "{"code":200,"data":[]model.Product}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/productList [get]
func (a *News) List(c *gin.Context) {
	service := model.NewNews()
	service.List(c)
	return
}

//Get 通过ID获取角色信息
//@Tags 后台产品模块
//@Summary 获取产品信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} model.Product "{"code":200,"data":model.Product}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/product/{id} [get]
func (a *News) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		service := model.NewNews()
		service.GetOne(id, c)
	}
	return
}

//Add 新增产品信息
//@Tags 后台产品模块
//@Summary 添加产品信息
//@Param token header string true "token"
//@Param name formData string true "名称"
//@Param title formData string true "标题"
//@Param version formData string true "型号"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/product [post]
func (a *News) Add(c *gin.Context) {
	service := model.NewNews()
	service.AddNews(c)
	return
}

//Edit 修改产品信息
//@Tags 后台产品模块
//@Summary 修改产品信息
//@Param token header string true "token"
//@Param name formData string true "名称"
//@Param title formData string true "标题"
//@Param version formData string true "型号"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/product/{id}/edit [put]
func (a *News) Edit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	service := model.NewNews()
	service.UpdateNews(id, c)
	return
}

//Delete 删除产品信息
//@Tags 后台产品模块
//@Summary 删除产品信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/product/{id} [delete]
func (a *News) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	service := model.NewNews()
	service.DeleteNews(id, c)
	return
}
