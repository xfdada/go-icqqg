package routers

import (
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/api/web"
	"gin-icqqg/config"
	"gin-icqqg/controller/admin"
	"gin-icqqg/controller/index"
	_ "gin-icqqg/docs"
	"gin-icqqg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resource/view/**/**/*") //加载模板文件
	r.Static("/asset", "public")            //静态资源
	r.StaticFS("/uploads", http.Dir(config.AppConfig.Upload.Path))
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Delims("{[", "]}")
	user := v1.User{}
	news := v1.News{}
	r.GET("/api/v1/captcha", v1.Captcha)
	r.GET("/api/v1/table", v1.GetTable)
	r.GET("/api/v1/table_info", v1.MyTable)
	r.POST("/api/v1/user/login", user.Login)
	r.POST("/api/v1/sendSms", v1.SendSms)
	r.GET("/index.html", index.Index)
	r.GET("/admin/helper", admin.Helpers)
	r.POST("/api/v1/user", user.AddUser)
	r.GET("/api/v1/news/:id", news.GetNewsById)
	r.DELETE("/api/v1/news/:id", news.DeleteNews)
	r.POST("/api/v1/news", news.AddNews)
	r.POST("/api/v1/upload", news.Upload)
	apiv1 := r.Group("api/v1")
	indexUser := &index.User{}
	apiv1.Use(middleware.Jwt(), middleware.Logger())
	{
		apiv1.GET("/index/user", indexUser.GetUserInfo)
		apiv1.POST("/user/update_avatar", user.UpdateAvatar)
		apiv1.PUT("/user", user.UpdateUserInfo)
		r.GET("/user/:id", user.GetUser)
	}
	webUser := web.NewWebUser()
	webMenu := web.NewAddMenu()
	webPermission := web.NewPermission()
	//webProduct := web.NewProduct()
	webRole := web.NewRole()
	r.GET("/api/web/getTree", webPermission.GetTree)
	adminLogin := web.NewAdminLogin()
	r.POST("/api/web/user/login", adminLogin.Login)
	admin := r.Group("api/web")
	admin.Use(middleware.AdminJwt())
	{
		admin.POST("/user", webUser.AddUser)
		admin.GET("/user", webUser.SelfInfo)
		admin.DELETE("/user/:username", webUser.DelUser)
		admin.GET("/userList", webUser.UserList)
		admin.PUT("/userIsOpen", webUser.IsOpen)
		admin.GET("/user/logout", adminLogin.LogOut)
		admin.GET("/menuList", webMenu.List)
		admin.POST("/menu", webMenu.AddMenu)
		admin.DELETE("/menu/:id", webMenu.DeleteMenu)
		admin.PUT("/menu/:id", webMenu.EditMenu)
		admin.GET("/menu/:id", webMenu.GetMenu)
		admin.GET("/menuTree", webMenu.GetTree)
		admin.GET("/permissionList", webPermission.List)
		admin.POST("/permission", webPermission.AddPermission)
		admin.DELETE("/permission/:id", webPermission.DeletePermission)
		admin.PUT("/permission/:id", webPermission.EditPermission)
		admin.GET("/permission/:id", webPermission.GetPermission)
		admin.GET("/permissionParent", webPermission.ParentList)
		admin.GET("/permissionTree", webPermission.GetTree)
		admin.Resource("/role", webRole)
		//admin.Resource("/product", webProduct)
	}

	return r
}
