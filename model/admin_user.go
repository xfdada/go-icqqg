package model

import (
	"gin-icqqg/config/response"
	"gin-icqqg/utils/jwt"
	mypassword "gin-icqqg/utils/password"
	"github.com/gin-gonic/gin"
)

/**
* AdminUser
* 模块主要是后端用户管理模块
 */

// AdminUser 后端用户结构体
type AdminUser struct {
	ID            int64  ` gorm:"comment:ID;type:int;size:10;primaryKey;autoIncrement;"    json:"id,omitempty"`
	UserName      string `gorm:"comment:账号名;type:varchar;size:120;uniqueIndex;column:username"   json:"user_name,omitempty"`
	Password      string `gorm:"comment:密码;type:varchar;size:191;column:password"  json:"-,omitempty"`
	Name          string `gorm:"comment:用户名;type:varchar;size:50;column:name"  json:"name,omitempty"`
	Avatar        string `gorm:"comment:头像;type:varchar;size:255;column:avatar"  json:"avatar,omitempty"`
	RememberToken string `gorm:"comment:免登录token;type:varchar;size:255;column:remember_token" json:"remember_token,omitempty"`
	CreatedAt     int64  ` gorm:"comment:创建时间;type:int;size:11;autoCreateTime"  json:"created_at,omitempty"`
	UpdatedAt     int64  `gorm:"comment:更新时间;type:int;size:11;autoUpdateTime"  json:"-,omitempty"`
	DeletedAt     int64  `gorm:"comment:删除时间;type:int;size:11;"  json:"-,omitempty"`
}

func NewAdminUser() *AdminUser {
	db.AutoMigrate(&AdminUser{})
	return &AdminUser{}
}

func (u *AdminUser) TableName() string {
	return "admin_user"
}

// GetUserInfo 获取用户信息 ，当前版本先获取用户基本信息，后续获取权限以及路由地址 通过用户的ID来或者用户名来
// Get
func (u *AdminUser) GetUserInfo() {

}

// GetSelfInfo 获取自身用户信息
// Get 可以通过token来实现
func (u *AdminUser) GetSelfInfo(username string, c *gin.Context) {
	res := db.Where("username = ?", username).First(u)
	if res.RowsAffected < 1 {
		c.JSON(500, gin.H{"code": 500, "msg": "未找到用户信息"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": u})
	}
	c.Abort()
}

// GetUserList 获取用户列表，可以分页
// Get
func (u *AdminUser) GetUserList() {

}

// AddUser 新增后端用户
// Post
func (u *AdminUser) AddUser(username, password, name string) (int64, error) {
	u.UserName = username
	u.Password, _ = mypassword.EncryptPassword(password)
	u.Name = name
	res := db.Save(&u)
	return res.RowsAffected, res.Error
}

// EditUser 编辑用户信息
//Put
func (u *AdminUser) EditUser() {

}

// ChangePassword 更改密码
//Put
func (u *AdminUser) ChangePassword() {

}

// UserLogin 用户登录
//Post
func (u *AdminUser) UserLogin(username, password string, c *gin.Context) {
	r := response.NewResponse(c)
	res := db.Where("username=? ", username).First(u)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.CountError)
	} else {
		if !mypassword.EqualsPassword(password, u.Password) {
			r.ErrorResp(response.CountError)
		} else {
			token, _ := jwt.GetAdminToken(u.UserName, u.ID)
			r.SuccessResp(map[string]interface{}{"code": 200, "token": token})
		}
	}
	c.Abort()
}
