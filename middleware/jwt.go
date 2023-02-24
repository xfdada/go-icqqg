package middleware

import (
	jwts "gin-icqqg/utils/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			eCode = 200
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			eCode = 300
		} else {

			_, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = 500
				default:
					eCode = 500
				}
			}
		}
		if eCode != 200 {
			c.Abort()
			return
		}
		c.Next()
	}
}
