package v1

import (
	"fmt"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type User struct{}

// GetUser
//@Tags 用户接口
//@Summary 获取用户详细信息
//@Param id  path int true "用户ID"
//@Param token header string true "用户token"
//@Produce json
//@Success 200   "成功"
// @Success 200 {object} model.User "成功"
// @Failure 400  "请求错误"
// @Failure 500  "内部错误"
//@Router /api/v1/user/{id} [get]
func (u User) GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	user := model.NewUser()
	data, row := user.GetUser(id)
	if row == 0 {
		c.JSON(500, gin.H{"msg": "用户不存在"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": data})
	c.Abort()
	return
}
