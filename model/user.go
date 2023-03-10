package model

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/utils/jwt"
	mypassword "gin-icqqg/utils/password"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	*Model
	Uuid      int64  ` json:"uuid" gorm:"primary_key;comment:uuid" `                                                      //uuid  自动生成
	EmailSing int    `json:"email_sing" gorm:"column:email_sing;comment:邮箱注册 1是 2否"`                                      //邮箱注册 1是 2否
	EmailBind int    `json:"email_bind" gorm:"column:email_bind;comment:邮箱绑定 1是 2否"`                                      //邮箱绑定 1是 2否
	Sex       int    `form:"sex" json:"sex" binding:"required;comment:性别 1男 2女 3未知"`                                      //性别 1男 2女 3未知
	UserName  string `form:"user_name" json:"user_name,omitempty" gorm:"column:user_name;comment:用户名" binding:"required"` //用户名
	Phone     string `form:"phone" json:"phone,omitempty" gorm:"comment:手机号" binding:"required"`                          //手机号
	Password  string `form:"password" json:"-,omitempty" gorm:"comment:密码" binding:"required"`                            //密码
	Email     string `form:"email" json:"email,omitempty" gorm:"comment:邮箱" binding:"required,email"`                     //邮箱
	Token     string `json:"token,omitempty" gorm:"comment:token"`                                                        //token
	OpenId    string `json:"-,omitempty" gorm:"comment:微信OpenId"`                                                         //微信OpenId
	Avatar    string `form:"avatar" json:"avatar,omitempty" gorm:"comment:头像"  binding:"required"`                        //头像
}
type UserUpdate struct {
	Sex      int    `form:"sex" json:"sex" binding:"required"`                                               //性别 1男 2女 3未知
	UserName string `form:"user_name" json:"user_name,omitempty" gorm:"column:user_name" binding:"required"` //用户名
	Phone    string `form:"phone" json:"phone,omitempty" gorm:"comment:手机号" binding:"required"`              //手机号
	Email    string `form:"email" json:"email,omitempty" binding:"required,email"`                           //邮箱
}

func (u User) TableName() string {
	return "table_user"
}

func NewUser() *User {
	db.AutoMigrate(&User{})
	return &User{}
}

func (u *User) GetUser(id string) (*User, int) {
	res := db.Where("id=?", id).First(&u)
	if res.Error != nil {
		fmt.Println(res.Error)
		return &User{}, 0
	}
	if res.RowsAffected <= 0 {
		return &User{}, 0
	}
	return u, 1
}

func (u *User) AddUser(c *gin.Context) {
	var users User
	r := response.NewResponse(c)
	err := c.ShouldBind(&users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	} else {
		passwords, _ := mypassword.EncryptPassword(users.Password)
		user := &User{
			Uuid:     config.Snow.GetId(),
			Sex:      users.Sex,
			UserName: users.UserName,
			Phone:    users.Phone,
			Password: passwords,
			Email:    users.Email,
			Avatar:   users.Avatar,
		}
		res := db.Create(user)
		if res.RowsAffected > 0 {
			r.SuccessResp("新增成功")
		} else {
			r.ErrorResp(response.AddError)
		}
	}

	c.Abort()
}

// Login  登录并且生成 token返回
func (u *User) Login(phone, password string, c *gin.Context) {
	r := response.NewResponse(c)
	res := db.Where("phone=? ", phone).First(u)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.CountError)
	} else {
		if !mypassword.EqualsPassword(password, u.Password) {
			r.ErrorResp(response.CountError)
		} else {
			token, _ := jwt.GenerateToken(u.Phone, strconv.FormatInt(u.Uuid, 10))
			r.SuccessResp(map[string]interface{}{"code": 200, "token": token})
		}
	}

	c.Abort()
}

// GetUserByUuid  前端通过jwt返回用户基本信息
func (u *User) GetUserByUuid(uuid string, c *gin.Context) {
	r := response.NewResponse(c)
	res := db.Where("uuid = ?", uuid).First(u)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.NotFoundError)
	} else {
		r.SuccessResp(u)
	}
	c.Abort()
}

// UpdateAvatar  前端通过jwt返回用户基本信息
func (u *User) UpdateAvatar(url, uuid string, c *gin.Context) {
	r := response.NewResponse(c)
	u.Avatar = url
	res := db.Where("uuid = ?", uuid).Updates(&u)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.UpdateError)
	} else {
		r.SuccessResp(map[string]string{"url": url})
	}
	c.Abort()
}

// UpdateUserInfo  更新用户基本信息
func (u *User) UpdateUserInfo(uuid string, c *gin.Context) {
	r := response.NewResponse(c)
	var userUpdate UserUpdate
	err := c.ShouldBind(&userUpdate)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		r.ErrorResp(response.ParamsError)
	} else {
		u.Phone = userUpdate.Phone
		u.Email = userUpdate.Email
		u.UserName = userUpdate.UserName
		u.Sex = userUpdate.Sex
		res := db.Where("uuid = ?", uuid).Updates(&u)
		if res.RowsAffected < 1 {
			r.ErrorResp(response.UpdateError)
		} else {
			r.SuccessResp("更新成功")
		}
	}

	c.Abort()
}
