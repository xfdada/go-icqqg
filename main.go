package main

import (
	"gin-icqqg/config"
	"gin-icqqg/routers"
)

// @title          ICQQG接囗文档
// @version        1.0
// @contact.name   Xfdada
// @contact.email  xiangfudada@163.com

func main() {
	router := routers.NewRouter()
	router.Run(config.AppConfig.Server.Port)
}
