package v1

import (
	"encoding/json"
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/utils/alisms"
	"gin-icqqg/utils/redis"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// SendSms @Tags 固定接口
//@Summary 发送短信
//@Param phone  formData string true " 手机号"
//@Produce json
//@Success 200   "成功"
// @Success 200 string string "成功"
// @Failure 400 string string "请求错误"
// @Failure 500 string string "内部错误"
//@Router /api/v1/sendSms [post]
func SendSms(c *gin.Context) {
	r := response.NewResponse(c)
	phone := c.PostForm("phone")
	code := fmt.Sprintf("%v", rand.Intn(8999)+1000)
	redis.Set("sms_"+phone, code, config.AppConfig.AlySms.Expiration*time.Second)
	param, _ := json.Marshal(map[string]string{"code": code})
	res := alisms.SendSmS(phone, string(param))
	if res {
		r.SuccessResp("发送成功")
	} else {
		r.ErrorResp(response.ServerError)
	}
	c.Abort()
	return
}
