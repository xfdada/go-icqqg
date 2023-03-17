package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ServerError   = NewCode(100, "服务器内部错误")
	ParamsError   = NewCode(101, "参数错误")
	NotFoundError = NewCode(102, "未找到结果")
	NoToken       = NewCode(103, "缺少token")
	TokenError    = NewCode(104, "鉴权失败，token错误")
	TokenTimeout  = NewCode(105, "鉴权失败，token超时无效，请重新生成")
	AddError      = NewCode(106, "添加失败")
	DeleteError   = NewCode(107, "删除失败")
	UpdateError   = NewCode(108, "更新失败")
	CodeError     = NewCode(109, "验证码错误，请重新输入")
	CountError    = NewCode(110, "账号或密码有误")
	NoCountError  = NewCode(111, "账号不存在")
	Logout        = NewCode(111, "已注销，请重新登录")
	Success       = NewCode(200, "成功")
)

type Response struct {
	Ctx *gin.Context
}

type Code struct {
	Code int
	Msg  string
}

func NewCode(code int, msg string) *Code {

	return &Code{
		Code: code,
		Msg:  msg,
	}
}

func (c *Code) Error() string {

	return fmt.Sprintf("错误码是:%d,错误信息是:%s", c.Code, c.Msg)
}

func (c *Code) Codes() int {
	return c.Code
}
func (c *Code) Msgs() string {
	return c.Msg
}

func (e *Code) StatusCode() int {
	switch e.Codes() {
	case Success.Codes():
		return http.StatusOK
	case ServerError.Codes():
		return http.StatusInternalServerError
	case NoToken.Codes():
		fallthrough
	case ParamsError.Codes():
		fallthrough
	case NotFoundError.Codes():
		fallthrough
	case NoCountError.Codes():
		fallthrough
	case CountError.Codes():
		return http.StatusBadRequest
	case TokenError.Codes():
		fallthrough
	case TokenTimeout.Codes():
		fallthrough
	case Logout.Codes():
		fallthrough
	case AddError.Codes():
		fallthrough
	case DeleteError.Codes():
		fallthrough
	case UpdateError.Codes():
		fallthrough
	case CodeError.Codes():
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//SuccessResp 成功返回
func (r *Response) SuccessResp(data interface{}) {
	r.Ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": data})
}

//ErrorResp 失败返回
func (r *Response) ErrorResp(c *Code) {
	response := gin.H{"code": c.Codes(), "msg": c.Msgs()}
	r.Ctx.JSON(c.StatusCode(), response)
}
