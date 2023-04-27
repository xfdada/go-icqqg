package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type ImMessage struct{}

func NewImMessage() *ImMessage {
	return &ImMessage{}
}

//List 数据列表
//@Tags 自动发送消息模块
//@Summary 消息列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.AutoMessage "{"code":200,"data":[]model.AutoMessage}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/imMessageList [get]
func (a *ImMessage) List(c *gin.Context) {
	service := model.NewImMessage()
	userId := c.Query("userId")
	service.GetUserList(userId, c)
	return
}
