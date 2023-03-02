package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	*Model
	Uuid      int64  `json:"uuid" gorm:"primary_key;"`                   //uuid
	EmailSing int    `json:"email_sing" gorm:"column:email_sing"`        //邮箱注册 1是 2否
	EmailBind int    `json:"email_bind" gorm:"column:email_bind"`        //邮箱绑定 1是 2否
	Sex       int    `json:"sex"`                                        //性别 1男 2女 3未知
	UserName  string `json:"user_name,omitempty" gorm:"column:username"` //用户名
	Phone     string `json:"phone,omitempty" gorm:"comment:手机号"`         //手机号
	Password  string `json:"-,omitempty"`                                //密码
	Email     string `json:"email,omitempty"`                            //邮箱
	Token     string `json:"token,omitempty"`                            //token
	OpenId    string `json:"-,omitempty"`                                //微信OpenId
	Avatar    string `json:"avatar,omitempty"`                           //头像
}

func (u User) TableName() string {
	return "table_user"
}

func NewUser() *User {
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
	var tmpParams = make(map[string]string)
	err2 := c.ShouldBind(&tmpParams)
	if err2 != nil {
		fmt.Println(err2)
	}
	user := &User{
		UserName: tmpParams["user_name"],
		Phone:    tmpParams["mobile"],
		Password: tmpParams["password"],
		Email:    tmpParams["email"],
	}
	res := db.Create(user)
	if res.RowsAffected > 0 {
		c.JSON(200, gin.H{"data": "success"})
	} else {
		c.JSON(500, gin.H{"msg": "用户新增失败"})
	}
	c.Abort()
}

func (u *User) Login(phone, password string) bool {
	res := db.Where("phone=? and password=?", phone, password).First(u)
	if res.RowsAffected < 1 {
		return false
	}
	return true
}
