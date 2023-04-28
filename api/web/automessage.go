package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type AutoMessage struct{}

func NewAutoMessage() *AutoMessage {
	return &AutoMessage{}
}

//List 数据列表
//@Tags 自动发送消息模块
//@Summary 消息列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.AutoMessage "{"code":200,"data":[]model.AutoMessage}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessageList [get]
func (a *AutoMessage) List(c *gin.Context) {
	service := model.NewAutoMessage()
	service.List(c)
	return
}

//Get 通过ID获取角色信息
//@Tags 自动发送消息模块
//@Summary 获取消息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} model.AutoMessage "{"code":200,"data":model.AutoMessage}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessage/{id} [get]
func (a *AutoMessage) Get(c *gin.Context) {
	service := model.NewAutoMessage()
	service.Get(c)
	return
}

//Add 新增产品信息
//@Tags 自动发送消息模块
//@Summary 新增消息
//@Param token header string true "token"
//@Param is_default formData string true "是否默认"
//@Param send_sort formData string true "发送顺序"
//@Param group_id formData string true "平台ID"
//@Param content formData string true "消息内容"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessage [post]
func (a *AutoMessage) Add(c *gin.Context) {
	service := model.NewAutoMessage()
	var add model.AddAutoMessage
	err := c.ShouldBind(&add)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err})
		c.Abort()
		return
	}
	service.Add(add, c)
	return
}

//Edit 修改产品信息
//@Tags 自动发送消息模块
//@Summary 修改消息
//@Param token header string true "token"
//@Param id path string true "ID"
//@Param is_default formData string true "是否默认"
//@Param send_sort formData string true "发送顺序"
//@Param group_id formData string true "平台ID"
//@Param content formData string true "消息内容"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessage/{id}/edit [put]
func (a *AutoMessage) Edit(c *gin.Context) {
	var add model.AddAutoMessage
	err := c.ShouldBind(&add)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": err})
		c.Abort()
		return
	}
	service := model.NewAutoMessage()
	service.Edit(add, c)
	return
}

//Delete 删除产品信息
//@Tags 自动发送消息模块
//@Summary 删除消息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessage/{id} [delete]
func (a *AutoMessage) Delete(c *gin.Context) {
	service := model.NewAutoMessage()
	service.Delete(c)
	return
}

//GetGroup 通过ID获取角色信息
//@Tags 自动发送消息模块
//@Summary 获取消息
//@Param token header string true "token"
//@Param group_id query string true "group_id"
//@Produce json
// @Success 200 {object} model.SendList "{"code":200,"data":[]model.AutoMessage}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/autoMessage/GetGroup [get]
func (a *AutoMessage) GetGroup(c *gin.Context) {
	service := model.NewAutoMessage()
	service.GetGroup(c)
	return
}
