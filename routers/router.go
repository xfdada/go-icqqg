package routers

import (
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/controller"
	_ "gin-icqqg/docs"
	"gin-icqqg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resource/view/**/*")
	r.Static("/asset", "public")
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := v1.User{}
	news := v1.News{}
	r.GET("/index", controller.Index)
	r.GET("/api/v1/get_token", v1.GetToken)
	r.POST("/api/v1/user", user.AddUser)
	r.GET("/api/v1/news/:id", news.GetNewsById)
	r.POST("/api/v1/news", news.AddNews)
	r.POST("/api/v1/upload", news.Upload)
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.Jwt(), middleware.Logger())
	{
		apiv1.GET("/user/:id", user.GetUser)

	}
	return r
}
