package routers

import (
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/api/web"
	"gin-icqqg/config"
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
	r.LoadHTMLGlob("resource/view/**/*") //加载模板文件
	r.Static("/asset", "public")         //静态资源
	r.StaticFS("/uploads", http.Dir(config.AppConfig.Upload.Path))
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := v1.User{}
	news := v1.News{}
	r.GET("/api/v1/captcha", v1.Captcha)
	r.GET("/api/v1/table", v1.GetTable)
	r.GET("/api/v1/table_info", v1.MyTable)
	r.POST("/api/v1/user/login", user.Login)
	r.POST("/api/v1/sendSms", v1.SendSms)
	r.GET("/index.html", index.Index)

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
	adminLogin := web.NewAdminLogin()
	r.POST("/api/web/user/login", adminLogin.Login)
	admin := r.Group("api/web")
	admin.Use(middleware.AdminJwt())
	{
		admin.POST("/user", webUser.AddUser)
		admin.GET("/user", webUser.SelfInfo)
		admin.GET("/user/logout", adminLogin.LogOut)
	}

	return r
}
