package middleware

import (
	"gin-icqqg/config/response"
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

			_, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = response.TokenTimeout
				default:
					eCode = response.TokenError
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
