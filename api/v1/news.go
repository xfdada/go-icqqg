package v1

import (
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/config/response"
	"gin-icqqg/model"
	"gin-icqqg/utils/upload"
	"github.com/gin-gonic/gin"
)

type News struct{}

//GetNewsById
//@Tags 新闻相关接口
//@Summary 通过ID获取新闻详细信息
//@Param id  path int true "新闻ID"
//@Produce json
//@Success 200   "成功"
//@Success 200 {object} model.News "成功"
//@Failure 400 {object} response.Code {"code": 101, "msg": "参数错误"} "请求错误"
//@Failure 500 string {"code": 102, "msg": "未找到结果"} "内部错误"
//@Router /api/v1/news/{id} [get]
func (n News) GetNewsById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	r := response.NewResponse(c)
	if id == "" {
		r.ErrorResp(response.ParamsError)
	} else {
		news := model.NewNews()
		data, num := news.GetById(id)
		if num > 0 {
			r.SuccessResp(data)
		} else {
			r.ErrorResp(response.NotFoundError)
		}
	}
	c.Abort()
	return
}

//AddNews
//@Tags 新闻相关接口
//@Summary 新增新闻
//@Param title formData string true "新闻标题"
//@Param desc formData string true "新闻描述"
//@Param thumb formData string true "新闻图"
//@Param seo_keyword formData string true "新闻关键词"
//@Param tags formData string true "新闻tag"
//@Param content formData string true "新闻内容"
//@Param cate_id formData string true "新闻类别"
//@Param zan formData int true "新闻点赞数"
//@Param hot formData int true "新闻热门"
//@Param see_num formData int true "新闻查看数"
//@Param type formData int true "新闻是否置顶"
//@Param show formData int true "新闻是否显示"
//@Produce json
//@Success 200   "成功"
//@Success 200 {object} response.Code "{"code": 200, "msg": "success"}"
//@Failure 400 {object} response.Code "{"code": 101, "msg": "参数错误"} 请求错误"
//@Failure 500 {object} response.Code "{"code": 102, "msg": "未找到结果"} 内部错误"
//@Router /api/v1/news [post]
func (n News) AddNews(c *gin.Context) {
	r := response.NewResponse(c)
	news := model.NewNews()
	if err := c.Bind(news); err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
		r.ErrorResp(response.ParamsError)
	} else {
		res := news.AddNews()
		if res > 0 {
			r.SuccessResp("success")
		} else {
			r.ErrorResp(response.AddError)
		}
	}
	c.Abort()
	return
}

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
		r.SuccessResp(url)
	}
	c.Abort()
	return
}

// DeleteNews
//@Tags 新闻相关接口
//@Summary 删除新闻
//@Param id path int true "新闻ID"
//@Produce json
//@Success 200   "成功"
//@Success 200 {object} response.Code "{"code": 200, "msg": "success"}"
//@Failure 400 {object} response.Code "{"code": 101, "msg": "参数错误"} 请求错误"
//@Failure 500 {object} response.Code "{"code": 102, "msg": "未找到结果"} 内部错误"
//@Router /api/v1/news/{id} [delete]
func (n *News) DeleteNews(c *gin.Context) {
	news := model.NewNews()
	news.DeleteNews(c)
	return
}
