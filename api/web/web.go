package web

type Web struct{}

type LoginUser struct {
	Id       int64
	UserName string
}

var Logins *LoginUser
