package v1

import (
	"fmt"
	jwts "gin-icqqg/utils/jwt"
	"github.com/gin-gonic/gin"
)

// GetToken @Tags 固定接口
//@Summary 获取token
//@Produce json
//@Success 200   "成功"
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
//@Router /api/v1/get_token [get]
func GetToken(c *gin.Context) {
	token, err := jwts.GenerateToken("xfdada", "go-admin") //这里可以自定义从数据库中取用户的账号和密码生成
	if err != nil {
		c.JSON(400, gin.H{"msg": "failed", "err": fmt.Sprintf("GenerateToken err:%v", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "token": token})
}
