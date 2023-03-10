package web

import (
	"fmt"
	"gin-icqqg/config/response"
	"gin-icqqg/model"
	"gin-icqqg/utils/captcha"
	"gin-icqqg/utils/redis"
	"github.com/gin-gonic/gin"
)

type AdminLogin struct {
	UserName     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
	CaptchaId    string `form:"captcha_id" binding:"required"`
	CaptchaValue string `form:"captcha_value" binding:"required"`
}

func NewAdminLogin() *AdminLogin {
	return new(AdminLogin)
}

//Login  登录接口
//@Tags 后端用户接口
//@Summary 用户登录
//@Param username  formData string true " 用户名"
//@Param password  formData string true " 密码"
//@Param captcha_id  formData string true " 验证码ID"
//@Param captcha_value  formData string true " 验证码"
//@Produce json
//@Success 200 {object} response.Code "成功"
//@Failure 400 {object} response.Code "请求错误"
//@Failure 500 {object} response.Code "内部错误"
//@Router /api/web/user/login [post]
func (u *AdminLogin) Login(c *gin.Context) {
	//首先验证码是否正确  false 返回验证码错误
	//在从数据库中查询用户是否存在
	//再次判断密码是否正确，如果都错误那么返回用户名或密码错误，
	var login AdminLogin
	r := response.NewResponse(c)
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(200, gin.H{"code": 401, "msg": fmt.Sprintf("%v", err)})
	}
	if !captcha.Verify(login.CaptchaId, login.CaptchaValue) {
		r.ErrorResp(response.CodeError)
	} else {
		user := model.NewAdminUser()
		user.UserLogin(login.UserName, login.Password, c)
	}
	return
}

//LogOut  用户退出
//@Tags 后端用户接口
//@Summary 用户注销
//@Param token  header string true " token"
//@Produce json
//@Success 200 {object} response.Code "{"code":200,"msg":"注销成功"}"
//@Failure 400 {object} response.Code "请求错误"
//@Failure 500 {object} response.Code "内部错误"
//@Router /api/web/user/logout [get]
func (u *AdminLogin) LogOut(c *gin.Context) {

	err := redis.Del(Logins.UserName)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "注销失败"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "注销成功"})
	}
	c.Abort()
	return
}
