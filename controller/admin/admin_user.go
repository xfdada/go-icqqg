package admin

type AdminUser struct {
	UserName string
	Id       int64
}

var AdminInfo *AdminUser

func GetAdminInfo() *AdminUser {
	if AdminInfo == nil {
		AdminInfo = new(AdminUser)
		return AdminInfo
	}
	return AdminInfo
}
