package web

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
	"strconv"
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

// UserList  获取用户列表
//@Tags 后端管理员模块
//@Summary 获取管理员列表
//@Param token header string true "token"
//@Param page query int true "页码"
//@Param pageSize query int true "每页数量"
//@Produce json
// @Success 200 {object} []model.AdminUser "{"code":200,"data":[]model.AdminUser}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/userList [get]
func (web *WebUser) UserList(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	tag := model.NewAdminUser()
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	tag.GetUserList(int(pageInt), int(pageSizeInt), c)
	c.Abort()
	return
}

// IsOpen  管理员是否启用
//@Tags 后端管理员模块
//@Summary 管理员是否启用
//@Param token header string true "token"
//@Param username formData string true "账号名"
//@Param is_open formData bool true "是否启用"
//@Produce json
// @Success 200 {object} []model.AdminUser "{"code":200,"data":[]model.AdminUser}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/userIsOpen [put]
func (web *WebUser) IsOpen(c *gin.Context) {
	isOpen := c.PostForm("is_open")
	username := c.PostForm("username")
	tag := model.NewAdminUser()
	open := 1
	if isOpen == "false" {
		open = 0
	}
	tag.UserIsOpen(username, open, c)
	c.Abort()
	return
}

//DelUser  删除管理员
//@Tags 后端管理员模块
//@Summary 删除管理员
//@Param token header string true "token"
//@Param username path string true "账号名"
//@Produce json
// @Success 200 {object} []model.AdminUser "{"code":200,"data":[]model.AdminUser}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/user/{username} [delete]
func (web *WebUser) DelUser(c *gin.Context) {
	username := c.Param("username")
	fmt.Println(username, "----------------")
	tag := model.NewAdminUser()
	tag.DelUser(username, c)
	c.Abort()
	return
}
