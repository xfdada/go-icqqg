package model

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/utils/jwt"
	mypassword "gin-icqqg/utils/password"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/**
* AdminUser
* 模块主要是后端用户管理模块
 */

// AdminUser 后端用户结构体
type AdminUser struct {
	ID            int64      ` gorm:"comment:ID;type:int;size:10;primaryKey;autoIncrement;"    json:"id,omitempty"`                //ID
	UserName      string     `gorm:"comment:账号名;type:varchar;size:120;column:username"   json:"user_name,omitempty"`               //用户名
	Password      string     `gorm:"comment:密码;type:varchar;size:191;column:password"  json:"-,omitempty"`                         //密码
	Name          string     `gorm:"comment:用户名;type:varchar;size:50;column:name"  json:"name,omitempty"`                          //名称
	Avatar        string     `gorm:"comment:头像;type:varchar;size:255;column:avatar"  json:"avatar,omitempty"`                      //头像
	ManageId      string     `gorm:"comment:管理员ID;type:varchar(255);column:manage_id"  json:"manage_id,omitempty"`                 //头像
	RememberToken string     `gorm:"comment:免登录token;type:varchar;size:255;column:remember_token" json:"remember_token,omitempty"` //免登录token
	IsOpen        int64      `gorm:"comment:是否启用;type:tinyint;size:3;column:is_open;default:1" json:"is_open,omitempty"`           //是否启用 1是 2否
	CreatedAt     int64      ` gorm:"comment:创建时间;type:int;size:11;autoCreateTime"  json:"created_at,omitempty"`
	UpdatedAt     int64      `gorm:"comment:更新时间;type:int;size:11;autoUpdateTime"  json:"-,omitempty"`
	DeletedAt     *LocalTime `gorm:"comment:删除时间;type:int;size:11;"  json:"-,omitempty"`
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
func (u *AdminUser) GetUserInfo(id string, c *gin.Context) {
	res := db.Where("id = ?", id).First(u)
	if res.RowsAffected < 1 {
		c.JSON(200, gin.H{"code": 500, "msg": "用户不存在"})
	} else {
		c.JSON(200, gin.H{"code": 200, "data": u})
	}
	c.Abort()
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

// GetUserList 获取用户列表，可以分页，返回数据有total，当前页码，当前条数，返回数据
// Get
func (u *AdminUser) GetUserList(page, pageSize int, c *gin.Context) {
	var List []*AdminUser
	dbpage := Paginate(page, pageSize)
	rows, err := dbpage(db).Model(&AdminUser{}).Rows()
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 500, "msg": "发送错误"})
	} else {
		for rows.Next() {
			db.ScanRows(rows, &List) //扫描到结构体中
		}
		var count int64
		var Res map[string]interface{}
		Res = make(map[string]interface{})
		Res["data"] = List
		Res["page"] = page
		db.Model(&AdminUser{}).Count(&count)
		Res["total"] = count
		c.JSON(200, gin.H{"code": 200, "data": Res})
	}
	c.Abort()
}

// AddUser 新增后端用户
// Post
func (u *AdminUser) AddUser(username, password, name string) (int64, error) {
	u.UserName = username
	u.Password, _ = mypassword.EncryptPassword(password)
	u.Name = name
	u.ManageId = uuid.New().String()
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

// UserLogin 用户登录  条件 账号启用
//Post
func (u *AdminUser) UserLogin(username, password string, c *gin.Context) {
	r := response.NewResponse(c)
	res := db.Where("username=? and is_open = 1", username).First(u)
	if res.RowsAffected < 1 {
		r.ErrorResp(response.NoCountError)
	} else {
		if !mypassword.EqualsPassword(password, u.Password) {
			r.ErrorResp(response.CountError)
		} else {
			fmt.Println(u)
			token, _ := jwt.GetAdminToken(u.ManageId, u.UserName, u.ID)
			tk := map[string]string{"token": token}
			Data := map[string]interface{}{"code": 200, "msg": "success", "data": tk}
			c.JSON(200, Data)
		}
	}
	c.Abort()
}

//UserIsOpen 用户是否启用
//put
func (u *AdminUser) UserIsOpen(username string, tag int, c *gin.Context) {
	err := db.Model(&AdminUser{}).Where("username = ?", username).Update("is_open", tag).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(500, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
}

//DelUser 删除用户保留信息
//delete
func (u *AdminUser) DelUser(username string, c *gin.Context) {
	err := db.Model(&AdminUser{}).Where("username = ?", username).Delete(&u).Error
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		c.JSON(500, gin.H{"code": 500, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success"})
	}
}
