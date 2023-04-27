package v1

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/utils/upload"
	"github.com/gin-gonic/gin"
)

type News struct{}

// Upload
//@Tags 图片上传接口
//@Summary 图片上传
// @Accept multipart/form-data
// @Param file formData file true "file"
//@Produce json
//@Success 200   "成功"
//@Success 200 {object} response.Code "{"code": 200, "msg": "success"}"
//@Failure 400 {object} response.Code "{"code": 101, "msg": "参数错误"} 请求错误"
//@Failure 500 {object} response.Code "{"code": 102, "msg": "未找到结果"} 内部错误"
//@Router /api/v1/upload [post]
func (n News) Upload(c *gin.Context) {
	r := response.NewResponse(c)
	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
	}
	url, err := upload.UploadFile(file, fileHead)

	if err != nil {
		r.ErrorResp(response.UpdateError)
	} else {

		c.JSON(200, gin.H{"link": "http://127.0.0.1:8080" + url})
	}
	c.Abort()
	return
}

// ImUpload
//@Tags 图片上传接口
//@Summary 图片上传
// @Accept multipart/form-data
// @Param file formData file true "file"
//@Produce json
//@Success 200   "成功"
//@Success 200 {object} response.Code "{"code": 200, "msg": "success"}"
//@Failure 400 {object} response.Code "{"code": 101, "msg": "参数错误"} 请求错误"
//@Failure 500 {object} response.Code "{"code": 102, "msg": "未找到结果"} 内部错误"
//@Router /api/v1/upload [post]
func (n News) ImUpload(c *gin.Context) {
	r := response.NewResponse(c)
	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
	}
	url, err := upload.UploadFile(file, fileHead)
	if err != nil {
		r.ErrorResp(response.UpdateError)
	} else {
		data := map[string]string{"src": "http://127.0.0.1:8080" + url}
		c.JSON(200, gin.H{"code": 200, "msg": "success", "data": data})
	}
	c.Abort()
	return
}
