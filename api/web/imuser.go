package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type ImUser struct {
}

func NewImUser() *ImUser {
	return &ImUser{}
}

//List 数据列表
//@Tags 后台产品模块
//@Summary 产品列表
//@Param token header string true "token"
//@Param page query string false "page"
//@Param pageNum query string false "pageNum"
//@Param start query string false "start"
//@Param end query string false "end"
//@Param manger query string false "manger"
//@Produce json
// @Success 200 {object} []model.ImUser "{"code":200,"data":[]model.ImUser}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/imUserList [get]
func (a *ImUser) List(c *gin.Context) {
	service := model.NewImUser()
	service.List(c)
	return
}
