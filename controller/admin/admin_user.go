package admin

type AdminUser struct {
	Id       int64
	UserName string
}

var AdminInfo *AdminUser

func GetAdminInfo() *AdminUser {
	if AdminInfo == nil {
		AdminInfo = new(AdminUser)
		return AdminInfo
	}
	return AdminInfo
}
