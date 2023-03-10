package middleware

import (
	"gin-icqqg/api/web"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	jwts "gin-icqqg/utils/jwt"
	"gin-icqqg/utils/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func AdminJwt() gin.HandlerFunc {
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

			res, err := jwts.ParseAdminToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = response.TokenTimeout
				default:
					eCode = response.TokenError
				}
			} else {
				web.Logins = &web.LoginUser{
					UserName: res.UserName,
					Id:       res.Id,
				}
				tokens, _ := redis.Get(res.UserName)
				if tokens != token {
					eCode = response.Logout
				} else {
					//判断 jwt 是否续签
					claims := res.StandardClaims
					if claims.ExpiresAt-time.Now().Unix() < config.AppConfig.Jwt.Renew {
						newToken, _ := jwts.GetAdminToken(res.UserName, res.Id)
						//将新的token注入到先赢头中
						c.Header("token", newToken)
					}
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
