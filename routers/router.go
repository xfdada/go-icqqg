package routers

import (
	v1 "gin-icqqg/api/v1"
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
	r.Use(middleware.Logger()).GET("/index", func(c *gin.Context) {
		c.String(200, "hahaha")
	})
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := v1.User{}
	r.GET("/api/v1/get_token", v1.GetToken)
	r.POST("/api/v1/user", user.AddUser)
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.Jwt(), middleware.Logger())
	{
		apiv1.GET("/user/:id", user.GetUser)

	}
	return r
}
