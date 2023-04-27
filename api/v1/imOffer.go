package v1

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type ImOffer struct {
}

func NewImOffer() *ImOffer {
	return &ImOffer{}
}

func (im *ImOffer) Offer(c *gin.Context) {
	offer := model.NewImOffer()
	offer.Add(c)
	return
}
