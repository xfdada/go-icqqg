package v1

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/controller/index"
	"gin-icqqg/model"
	"gin-icqqg/utils/captcha"
	"gin-icqqg/utils/upload"
	"github.com/gin-gonic/gin"
)

type User struct{}
type UserLogin struct {
	Phone        string `form:"phone" binding:"required"`
	Password     string `form:"password" binding:"required"`
	CaptchaId    string `form:"captcha_id" binding:"required"`
	CaptchaValue string `form:"captcha_value" binding:"required"`
}

type UserEdit struct {
	Sex      int    `form:"sex" json:"sex" binding:"required"`                                               //性别 1男 2女 3未知
	UserName string `form:"user_name" json:"user_name,omitempty" gorm:"column:user_name" binding:"required"` //用户名
	Phone    string `form:"phone" json:"phone,omitempty" gorm:"comment:手机号" binding:"required"`              //手机号
	Password string `form:"password" json:"-,omitempty" binding:"required"`                                  //密码
	Email    string `form:"email" json:"email,omitempty" binding:"required,email"`                           //邮箱
}

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
//@Param phone  formData string true " 手机号"
//@Param password  formData string true " 密码"
//@Param email  formData string true " 邮箱"
//@Param avatar  formData string true " 头像"
//@Param sex  formData integer true " 性别"
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
	var login UserLogin
	r := response.NewResponse(c)
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(200, gin.H{"code": 401, "msg": fmt.Sprintf("%v", err)})
	}
	if !captcha.Verify(login.CaptchaId, login.CaptchaValue) {
		r.ErrorResp(response.CodeError)
	} else {
		user := model.NewUser()
		user.Login(login.Phone, login.Password, c)
	}
	return
}

// UpdateAvatar
//@Tags 用户接口
//@Summary 用户更换头像
// @Accept multipart/form-data
// @Param file formData file true "file"
//@Param token  header string true "用户token"
//@Produce json
// @Success 200 {object} response.Code  "{"code":200,"url":""}成功"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/v1/user/update_avatar [post]
func (u *User) UpdateAvatar(c *gin.Context) {
	r := response.NewResponse(c)
	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
	}
	url, err := upload.UploadFile(file, fileHead)
	if err != nil {
		r.ErrorResp(response.ServerError)
		c.Abort()
	} else {
		uuid := index.Userinfo.Uuid
		user := model.NewUser()
		user.UpdateAvatar(url, uuid, c)
	}

	return

}

// UpdateUserInfo
//@Tags 用户接口
//@Summary 用户更新信息
// @Accept multipart/form-data
//@Param user_name  formData string true " 用户名"
//@Param sex  formData string true " 性别"
//@Param phone  formData string true " 手机号"
//@Param email  formData string true " 邮箱"
//@Param token  header string true "用户token"
//@Produce json
// @Success 200 {object} response.Code  "{"code":200,"data":"更新成功"}成功"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/v1/user [put]
func (u *User) UpdateUserInfo(c *gin.Context) {

	uuid := index.Userinfo.Uuid
	user := model.NewUser()
	user.UpdateUserInfo(uuid, c)

	return

}
