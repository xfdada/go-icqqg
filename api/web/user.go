package web

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type WebUser struct{}

type AddUser struct {
	UserName string `form:"username" gorm:"comment:账号名;type:varchar;size:120;uniqueIndex;column:username" binding:"required"   json:"user_name,omitempty"`
	Password string `form:"password" gorm:"comment:密码;type:varchar;size:191;column:password" binding:"required"  json:"-,omitempty"`
	Name     string `form:"name" gorm:"comment:用户名;type:varchar;size:50;column:name" binding:"required" json:"name,omitempty"`
	//后面还有角色id
}

func NewWebUser() *WebUser {
	return new(WebUser)
}

// AddUser  新增管理员
//@Tags 后端管理员模块
//@Summary 新增管理员
//@Param username  formData string true " 用户名"
//@Param password  formData string true " 密码"
//@Param name  formData string true " 姓名"
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"新增成功"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/user [post]
func (web *WebUser) AddUser(c *gin.Context) {
	var user AddUser
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(401, gin.H{"code": 401, "msg": fmt.Sprintf("%v", err)})
	} else {
		tag := model.NewAdminUser()
		res, err := tag.AddUser(user.UserName, user.Password, user.Name)
		if res >= 1 && err == nil {
			c.JSON(200, gin.H{"code": 200, "msg": "新增成功"})
		} else {
			config.ErrorLog(fmt.Sprintf("%v", err))
			c.JSON(402, gin.H{"code": 402, "msg": "新增失败"})
		}
	}
	c.Abort()
	return
}

// SelfInfo  获取自身信息
//@Tags 后端管理员模块
//@Summary 获取自身信息
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} model.AdminUser "{"code":200,"data":model.AdminUser}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/user [get]
func (web *WebUser) SelfInfo(c *gin.Context) {
	tag := model.NewAdminUser()
	tag.GetSelfInfo(Logins.UserName, c)
	return
}
