package routers

import (
	"fmt"
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/config"
	"gin-icqqg/controller/index"
	_ "gin-icqqg/docs"
	"gin-icqqg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
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
	r.GET("/", func(context *gin.Context) {
	loop:
		for {
			select {
			case <-time.After(20 * time.Second):
				fmt.Println("down")
				goto loop
			default:
				time.Sleep(time.Second * 2)
				fmt.Println("程序正在运行中")
			}
		}

	})
	r.GET("/api/v1/captcha", user.Captcha)
	r.POST("/api/v1/user/login", user.Login)
	r.POST("/api/v1/sendSms", v1.SendSms)
	r.GET("/index.html", index.Index)
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
