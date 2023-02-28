package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{"title": "go admin index", "h1": "这是模板"})
}
