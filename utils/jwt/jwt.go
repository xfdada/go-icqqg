package jwt

import (
	"gin-icqqg/config"
	"gin-icqqg/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Phone     string `json:"phone,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	AppKey    string `json:"_"`
	AppSecret string `json:"_"`
	jwt.StandardClaims
}

// GetJWTSecret 获取密钥
func GetJWTSecret() []byte {
	return []byte(config.AppConfig.Jwt.AppSecret)
}

// GenerateToken  生成token
func GenerateToken(appkey, appsecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Second * config.AppConfig.Jwt.Expire)
	claims := Claims{
		AppKey:    utils.EncodeMD5(appkey),
		AppSecret: utils.EncodeMD5(appsecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppConfig.Jwt.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
