package index

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type User struct {
	Uuid      int64  ` json:"uuid" `                                                         //uuid  自动生成
	EmailSing int    `json:"email_sing" gorm:"column:email_sing"`                            //邮箱注册 1是 2否
	EmailBind int    `json:"email_bind" gorm:"column:email_bind"`                            //邮箱绑定 1是 2否
	Sex       int    `json:"sex" binding:"required"`                                         //性别 1男 2女 3未知
	UserName  string `json:"user_name,omitempty" gorm:"column:user_name" binding:"required"` //用户名
	Phone     string `json:"phone,omitempty" gorm:"comment:手机号" binding:"required"`          //手机号
	Password  string `json:"-,omitempty" binding:"required"`                                 //密码
	Email     string `form:"email" json:"email,omitempty" binding:"required,email"`          //邮箱
	Token     string `json:"-,omitempty"`                                                    //token
	OpenId    string `json:"-,omitempty" gorm:"comment:微信OpenId"`                            //微信OpenId
	Avatar    string `form:"avatar" json:"avatar,omitempty"  binding:"required"`             //头像
}

func (u *User) GetUserInfo(c *gin.Context) {
	user := model.NewUser()
	user.GetUserByUuid(Userinfo.Uuid, c)
	return
}
