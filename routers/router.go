package routers

import (
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/api/web"
	"gin-icqqg/config"
	"gin-icqqg/controller/admin"
	"gin-icqqg/controller/index"
	_ "gin-icqqg/docs"
	im2 "gin-icqqg/im"
	"gin-icqqg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	configs := cors.DefaultConfig()
	configs.AllowOrigins = []string{"*"}
	configs.AllowHeaders = []string{"token"}
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upGrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	r.Use(gin.Logger(), cors.New(configs))
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resource/view/**/*") //加载模板文件
	r.Static("/asset", "public")         //静态资源
	r.StaticFS("/uploads", http.Dir(config.AppConfig.Upload.Path))
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Delims("{[", "]}")
	user := v1.User{}
	news := v1.News{}
	im := im2.NewIm(upGrader)
	imUser := v1.NewImUser()
	imOffer := v1.NewImOffer()
	r.GET("/api/v1/getList", imUser.List)
	r.GET("/ws", im.Toke)
	r.POST("/api/v1/offer", imOffer.Offer)
	r.GET("/api/v1/captcha", v1.Captcha)
	r.GET("/api/v1/table", v1.GetTable)
	r.GET("/api/v1/table_info", v1.MyTable)
	r.POST("/api/v1/user/login", user.Login)
	r.POST("/api/v1/sendSms", v1.SendSms)
	r.GET("/index.html", index.Index)
	r.GET("/admin/helper", admin.Helpers)
	r.POST("/api/v1/user", user.AddUser)
	r.POST("/api/v1/upload", news.Upload)
	r.POST("/api/v1/imUpload", news.ImUpload)
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
	webProduct := web.NewProduct()
	webRole := web.NewRole()
	webNews := web.NewNews()
	webAuto := web.NewAutoMessage()
	ImUser := web.NewImUser()
	ImMessage := web.NewImMessage()
	ImCode := web.NewImCode()
	webImOffer := web.NewImOffer()
	webFlow := web.NewFlow()
	r.GET("/api/web/flowList", webFlow.List)
	r.GET("/api/web/flowHour", webFlow.GetByHour)
	r.Resource("/api/web/imCode", ImCode)
	r.Resource("/api/web/autoMessage", webAuto)
	r.GET("/api/web/autoMessage/GetGroup", webAuto.GetGroup)
	r.GET("/api/web/imUserList", ImUser.List)
	r.GET("/api/web/imMessageList", ImMessage.List)
	r.Resource("/api/web/news", webNews)
	r.GET("/api/web/getTree", webPermission.GetTree)
	r.GET("/api/web/menuTree", webMenu.GetTree)
	r.GET("/api/web/imOfferList", webImOffer.List)
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

		admin.GET("/permissionList", webPermission.List)
		admin.POST("/permission", webPermission.AddPermission)
		admin.DELETE("/permission/:id", webPermission.DeletePermission)
		admin.PUT("/permission/:id", webPermission.EditPermission)
		admin.GET("/permission/:id", webPermission.GetPermission)
		admin.GET("/permissionParent", webPermission.ParentList)
		admin.GET("/permissionTree", webPermission.GetTree)

		admin.Resource("/role", webRole)
		admin.Resource("/menu", webMenu)
		admin.Resource("/product", webProduct)

	}

	return r
}
