package model

import "fmt"

type User struct {
	*Model
	UserName string `json:"user_name" gorm:"column:username" json:"user_name,omitempty"` //用户名
	Mobile   string `json:"mobile,omitempty" gorm:"comment:手机号"`                         //手机号
	Password string `json:"password,omitempty"`                                          //密码
	Email    string `json:"email,omitempty"`                                             //邮箱
	Token    string `json:"token,omitempty"`                                             //token
}

func (u User) TableName() string {
	return "ssf_user"
}

func NewUser() User {
	return User{}
}

func (u User) GetUser(id string) (User, int) {
	res := db.Where("id=?", id).First(&u)
	fmt.Println(res)
	if res.Error != nil {
		fmt.Println(res.Error)
		return User{}, 0
	}
	if res.RowsAffected <= 0 {
		return User{}, 0
	}
	return u, 1
}
