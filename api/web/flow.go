package web

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type Flow struct {
}

func NewFlow() *Flow {
	return &Flow{}
}

func (f *Flow) List(c *gin.Context) {
	fl := model.NewFlow()
	fl.List(c)
	return
}

func (f *Flow) GetByHour(c *gin.Context) {
	fl := model.NewFlow()
	fl.GetByHour(c)
	return
}
