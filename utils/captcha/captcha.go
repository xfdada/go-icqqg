package captcha

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/utils/loggers"
	"gin-icqqg/utils/redis"
	"github.com/mojocn/base64Captcha"
	"time"
)

type Store struct {
	Expiration time.Duration
	PreKey     string
}

var captcha = config.AppConfig.CaptCha

func NewDefaultRedisStore() base64Captcha.Store {
	return &Store{
		Expiration: time.Second * captcha.Expiration,
		PreKey:     captcha.PreKey,
	}
}

func (rs *Store) Set(id string, value string) error {
	err := redis.Set(rs.PreKey+id, value, rs.Expiration)
	if err != nil {
		loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
		return err
	}
	return nil

}
func (rs *Store) Get(key string, clear bool) string {
	val, err := redis.Get(key)
	if err != nil {
		loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
		return ""
	}
	if clear {
		err := redis.Del(key)
		if err != nil {
			loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
			return ""
		}
	}
	return val
}

func (rs *Store) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

// GetCaptcha 生成base64验证码
func GetCaptcha() (string, string) {
	driver := base64Captcha.NewDriverDigit(captcha.Height, captcha.Width, captcha.Length, captcha.MaxSkew, captcha.DotCount)
	// 生成base64图片
	store := base64Captcha.DefaultMemStore
	if captcha.UseRedis {
		store = NewDefaultRedisStore()
	}
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		loggers.Logs(fmt.Sprintf("Register GetCaptchaPhoto get base64Captcha has err:%v", err))
		return "", ""
	}
	return id, b64s
}

// Verify 验证是否正确
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	store := base64Captcha.DefaultMemStore
	if captcha.UseRedis {
		store = NewDefaultRedisStore()
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
