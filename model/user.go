package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	*Model
	UserName string `json:"user_name,omitempty" gorm:"column:username"` //用户名
	Mobile   string `json:"mobile,omitempty" gorm:"comment:手机号"`        //手机号
	Password string `json:"password,omitempty"`                         //密码
	Email    string `json:"email,omitempty"`                            //邮箱
	Token    string `json:"token,omitempty"`                            //token
}

func (u User) TableName() string {
	return "ssf_user"
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
		Mobile:   tmpParams["mobile"],
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
