package jwt

import (
	"gin-icqqg/config"
	"gin-icqqg/utils"
	"gin-icqqg/utils/redis"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Phone     string `json:"phone,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
	AppKey    string `json:"_"`
	AppSecret string `json:"_"`
	jwt.StandardClaims
}

type AdminClaims struct {
	UserName  string `json:"user_name,omitempty"`
	Id        int64  `json:"id,omitempty"`
	AppKey    string `json:"_"`
	AppSecret string `json:"_"`
	jwt.StandardClaims
}

// GetJWTSecret 获取密钥
func GetJWTSecret() []byte {
	return []byte(config.AppConfig.Jwt.AppSecret)
}

// GenerateToken  生成token
func GenerateToken(phone, uuid string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Second * config.AppConfig.Jwt.Expire)
	claims := Claims{
		Phone:     phone,
		Uuid:      uuid,
		AppKey:    utils.EncodeMD5(config.AppConfig.Jwt.AppKey),
		AppSecret: utils.EncodeMD5(config.AppConfig.Jwt.AppSecret),
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

// GetAdminToken  生成token
func GetAdminToken(userName string, id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Second * config.AppConfig.Jwt.Expire)
	claims := AdminClaims{
		UserName:  userName,
		Id:        id,
		AppKey:    utils.EncodeMD5(config.AppConfig.Jwt.AppKey),
		AppSecret: utils.EncodeMD5(config.AppConfig.Jwt.AppSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppConfig.Jwt.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	redis.Set(userName, token, time.Second*config.AppConfig.Jwt.Expire)
	return token, err
}

// ParseAdminToken  解析
func ParseAdminToken(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
