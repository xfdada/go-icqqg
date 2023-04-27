package v1

import (
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type ImUser struct {
}

func NewImUser() *ImUser {
	return &ImUser{}
}

func (im *ImUser) List(c *gin.Context) {
	user := model.NewImUser()
	user.GetFriendList("10086", c)
	return
}
