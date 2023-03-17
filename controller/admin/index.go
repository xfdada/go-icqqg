package admin

import "github.com/gin-gonic/gin"

//Index 代码生成器
func Index(c *gin.Context) {
	c.HTML(200, "admin/comment/public.html", nil)
}

func HomeIndex(c *gin.Context) {
	c.HTML(200, "admin/home/index.html", nil)
}
