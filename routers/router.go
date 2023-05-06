package routers

import (
	v1 "gin-icqqg/api/v1"
	"gin-icqqg/api/web"
	"gin-icqqg/config"
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

var router *gin.Engine

func NewRouter() *gin.Engine {
	router = gin.New()
	gin.SetMode(config.AppConfig.Server.Model)
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
	router.Use(gin.Logger(), cors.New(configs))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("resource/view/**/*") //加载模板文件
	router.Static("/asset", "public")         //静态资源
	router.StaticFS("/uploads", http.Dir(config.AppConfig.Upload.Path))
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Delims("{[", "]}")
	user := v1.User{}
	news := v1.News{}
	im := im2.NewIm(upGrader)
	imUser := v1.NewImUser()
	imOffer := v1.NewImOffer()
	router.GET("/api/v1/getList", imUser.List)
	router.GET("/ws", im.Toke)
	router.POST("/api/v1/offer", imOffer.Offer)
	router.GET("/api/web/captcha", v1.Captcha)
	//r.GET("/api/v1/table", v1.GetTable)
	//r.GET("/api/v1/table_info", v1.MyTable)
	router.POST("/api/v1/user/login", user.Login)
	router.POST("/api/v1/sendSms", v1.SendSms)
	router.GET("/index.html", index.Index)
	router.POST("/api/v1/user", user.AddUser)
	router.POST("/api/v1/upload", news.Upload)
	router.POST("/api/v1/imUpload", news.ImUpload)
	apiv1 := router.Group("api/v1")
	indexUser := &index.User{}
	apiv1.Use(middleware.Jwt(), middleware.Logger())
	{
		apiv1.GET("/index/user", indexUser.GetUserInfo)
		apiv1.POST("/user/update_avatar", user.UpdateAvatar)
		apiv1.PUT("/user", user.UpdateUserInfo)
		apiv1.GET("/user/:id", user.GetUser)
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

	adminLogin := web.NewAdminLogin()
	router.POST("/api/web/user/login", adminLogin.Login)
	admin := router.Group("api/web")
	admin.Use(middleware.AdminJwt())
	{
		admin.POST("/user", webUser.AddUser)
		admin.GET("/user", webUser.SelfInfo)
		admin.DELETE("/user/:username", webUser.DelUser)
		admin.GET("/userList", webUser.UserList)
		admin.PUT("/userIsOpen", webUser.IsOpen)
		admin.GET("/user/logout", adminLogin.LogOut)
		admin.GET("/getFriendList", ImUser.FriendList)
		admin.GET("/permissionList", webPermission.List)
		admin.POST("/permission", webPermission.AddPermission)
		admin.DELETE("/permission/:id", webPermission.DeletePermission)
		admin.PUT("/permission/:id", webPermission.EditPermission)
		admin.GET("/permission/:id", webPermission.GetPermission)
		admin.GET("/permissionParent", webPermission.ParentList)
		admin.GET("/permissionTree", webPermission.GetTree)
		admin.GET("/flowList", webFlow.List)
		admin.GET("/flowHour", webFlow.GetByHour)
		admin.GET("/autoMessage/GetGroup", webAuto.GetGroup)
		admin.GET("/imUserList", ImUser.List)
		admin.GET("/imMessageList", ImMessage.List)
		admin.GET("/getTree", webPermission.GetTree)
		admin.GET("/menuTree", webMenu.GetTree)
		admin.GET("/imOfferList", webImOffer.List)

		admin.GET("/newsList", webNews.List)
		admin.GET("/roleList", webRole.List)
		admin.GET("/menuList", webMenu.List)
		admin.GET("/productList", webProduct.List)
		admin.GET("/imCodeList", ImCode.List)
		admin.GET("/autoMessageList", webAuto.List)

		admin.POST("/news", webNews.Add)
		admin.POST("/role", webRole.Add)
		admin.POST("/menu", webMenu.Add)
		admin.POST("/product", webProduct.Add)
		admin.POST("/imCode", ImCode.Add)
		admin.POST("/autoMessage", webAuto.Add)

		admin.DELETE("/news/:id", webNews.Delete)
		admin.DELETE("/role/:id", webRole.Delete)
		admin.DELETE("/menu/:id", webMenu.Delete)
		admin.DELETE("/product/:id", webProduct.Delete)
		admin.DELETE("/imCode/:id", ImCode.Delete)
		admin.DELETE("/autoMessage/:id", webAuto.Delete)

		admin.PUT("/news/:id/edit", webNews.Edit)
		admin.PUT("/role/:id/edit", webRole.Edit)
		admin.PUT("/menu/:id/edit", webMenu.Edit)
		admin.PUT("/product/:id/edit", webProduct.Edit)
		admin.PUT("/imCode/:id/edit", ImCode.Edit)
		admin.PUT("/autoMessage/:id/edit", webAuto.Edit)

		admin.GET("/news/:id", webNews.Get)
		admin.GET("/role/:id", webRole.Get)
		admin.GET("/menu/:id", webMenu.Get)
		admin.GET("/product/:id", webProduct.Get)
		admin.GET("/imCode/:id", ImCode.Get)
		admin.GET("/autoMessage/:id", webAuto.Get)

	}

	return router
}
