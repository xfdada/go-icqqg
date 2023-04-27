package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type ImOffer struct {
}

func NewImOffer() *ImOffer {
	return &ImOffer{}
}

func (im *ImOffer) List(c *gin.Context) {
	list := model.NewImOffer()
	list.List(c)
	return
}
