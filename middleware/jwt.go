package middleware

import (
	"gin-icqqg/config/response"
	"gin-icqqg/controller/index"
	jwts "gin-icqqg/utils/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			eCode = response.Success
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			eCode = response.NoToken
		} else {

			res, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = response.TokenTimeout
				default:
					eCode = response.TokenError
				}
			} else {
				index.Userinfo = &index.UserInfo{
					Phone: res.Phone,
					Uuid:  res.Uuid,
				}
			}
		}
		if eCode != response.Success {
			//返回错误请求
			r := response.NewResponse(c)
			r.ErrorResp(eCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
