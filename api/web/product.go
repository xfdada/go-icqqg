package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type Product struct {
    Name       string     `form:"name" json:"name,omitempty"`    //名称
    Title       string     `form:"title" json:"title,omitempty"`    //标题
    Version       string     `form:"version" json:"version,omitempty"`    //型号
}

func NewProduct() *Product {
	return &Product{}
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
func (a *Product) List(c *gin.Context) {
	service := model.NewProduct()
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
func (a *Product) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		service := model.NewProduct()
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
func (a *Product) Add(c *gin.Context) {
    service := model.NewProduct()
    service.AddProduct(c)
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
func (a *Product) Edit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
    service := model.NewProduct()
    service.UpdateProduct(id,c)
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
func (a *Product) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	service := model.NewProduct()
	service.DeleteProduct(id, c)
	return
}
