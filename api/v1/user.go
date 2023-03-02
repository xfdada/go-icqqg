package v1

import (
	"gin-icqqg/config/response"
	"gin-icqqg/model"
	"gin-icqqg/utils/captcha"
	"github.com/gin-gonic/gin"
)

type User struct{}

// GetUser
//@Tags 用户接口
//@Summary 获取用户详细信息
//@Param id  path int true "用户ID"
//@Param token header string true "用户token"
//@Produce json
//@Success 200   "成功"
// @Success 200 {object} model.User "成功"
// @Failure 400  "请求错误"
// @Failure 500  "内部错误"
//@Router /api/v1/user/{id} [get]
func (u User) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := model.NewUser()
	data, row := user.GetUser(id)
	if row == 0 {
		c.JSON(500, gin.H{"msg": "用户不存在"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": data})
	c.Abort()
	return
}

// AddUser
//@Tags 用户接口
//@Summary 新增用户
//@Param user_name  formData string true " 用户名"
//@Param mobile  formData string true " 手机号"
//@Param password  formData string true " 密码"
//@Param email  formData string true " 邮箱"
//@Produce json
// @Success 200  "成功"
// @Failure 400  "请求错误"
// @Failure 500  "内部错误"
//@Router /api/v1/user [post]
func (u User) AddUser(c *gin.Context) {
	user := model.NewUser()
	user.AddUser(c)

	return
}

// Login  登录接口
//@Tags 用户接口
//@Summary 用户登录
//@Param phone  formData string true " 手机号"
//@Param password  formData string true " 密码"
//@Param captcha_id  formData string true " 验证码ID"
//@Param captcha_value  formData string true " 验证码"
//@Produce json
// @Success 200 {object} response.Code "成功"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/v1/user/login [post]
func (u User) Login(c *gin.Context) {
	//首先验证码是否正确  false 返回验证码错误
	//在从数据库中查询用户是否存在
	//再次判断密码是否正确，如果都错误那么返回用户名或密码错误，
	r := response.NewResponse(c)
	cId := c.PostForm("captcha_id")
	cValue := c.PostForm("captcha_value")
	if cId == "" || cValue == "" || !captcha.Verify(cId, cValue) {
		r.ErrorResp(response.CodeError)
	} else {
		phone := c.PostForm("phone")
		password := c.PostForm("password")
		if phone == "" || password == "" {
			r.ErrorResp(response.ParamsError)
		} else {
			user := model.NewUser()
			if user.Login(phone, password) {
				//这里可以生成token
				token := UserToken(phone)
				r.SuccessResp(map[string]string{"msg": "登录成功", "token": token}) //
			} else {
				r.ErrorResp(response.NotFoundError)
			}
		}
	}
	c.Abort()
	return
}

func (u User) Captcha(c *gin.Context) {
	r := response.NewResponse(c)
	id, url := captcha.GetCaptcha()
	r.SuccessResp(map[string]string{"id": id, "url": url})
}
