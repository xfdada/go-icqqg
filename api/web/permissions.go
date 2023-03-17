package web

import (
	"fmt"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type Permission struct {
	Name       string `form:"name" json:"name,omitempty" binding:"required"`           //权限英文名
	Tag        string `form:"tag" json:"tag,omitempty" binding:"required"`             //权限名称
	HttpMethod string `form:"http_method" json:"http_method,omitempty"`                //请求方法  空为任何方法 ，GET，POST，PUT，DELETE，PATCH，OPTIONS
	HttpPath   string `form:"http_path" json:"http_path,omitempty" binding:"required"` //请求路径
	ParentId   int64  `form:"parent_id" json:"parent_id,omitempty" `                   //父级ID
}
type EditPermission struct {
	Name       string `form:"name" json:"name,omitempty"`               //权限英文名
	Tag        string `form:"tag" json:"tag,omitempty"`                 //权限名称
	HttpMethod string `form:"http_method" json:"http_method,omitempty"` //请求方法  空为任何方法 ，GET，POST，PUT，DELETE，PATCH，OPTIONS
	HttpPath   string `form:"http_path" json:"http_path,omitempty"`     //请求路径
	ParentId   int64  `form:"parent_id" json:"parent_id,omitempty"`     //父级ID
}

func NewPermission() *Permission {

	return &Permission{}
}

//List 数据列表
//@Tags 后端权限模块
//@Summary 权限列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.Permissions "{"code":200,"data":[]model.Permission}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permissionList [get]
func (a *Permission) List(c *gin.Context) {
	menu := model.NewPermissions()
	menu.List(c)
	return
}

//ParentList 数据列表
//@Tags 后端权限模块
//@Summary 顶级权限列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.Permissions "{"code":200,"data":[]model.Permission}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permissionList [get]
func (a *Permission) ParentList(c *gin.Context) {
	menu := model.NewPermissions()
	menu.GetParent(c)
	return
}

//GetPermission 通过ID获取权限信息
//@Tags 后端权限模块
//@Summary 获取权限信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} model.Permissions "{"code":200,"data":model.Permission}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permission/{id} [get]
func (a *Permission) GetPermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		menu := model.NewPermissions()
		menu.GetOne(id, c)
	}
	return
}

//AddPermission 新增权限信息
//@Tags 后端权限模块
//@Summary 添加权限信息
//@Param token header string true "token"
//@Param name formData string true "权限名"
//@Param tag formData string true "权限中文名"
//@Param http_path formData string true "请求路径"
//@Param http_method formData string false "空为any 任何方法，get,put,post"
//@Param parent_id formData int false "父级ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permission [post]
func (a *Permission) AddPermission(c *gin.Context) {
	params := map[string]interface{}{}
	err := c.ShouldBind(&a)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["name"] = a.Name
		params["tag"] = a.Tag
		params["http_path"] = a.HttpPath
		params["http_method"] = a.HttpMethod
		params["parent_id"] = a.ParentId
		menu := model.NewPermissions()
		menu.AddPermissions(params, c)
	}
	return
}

//EditPermission 修改权限信息
//@Tags 后端权限模块
//@Summary 修改权限信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Param name formData string false "权限名"
//@Param tag formData string false "权限中文名"
//@Param http_path formData string false "请求路径"
//@Param http_method formData string false "空为any 任何方法，get,put,post"
//@Param parent_id formData int false "父级ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permission/{id} [put]
func (a *Permission) EditPermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	params := map[string]interface{}{}
	var edit EditPermission
	err := c.ShouldBind(&edit)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["name"] = edit.Name
		params["parent_id"] = edit.ParentId
		params["http_path"] = edit.HttpPath
		params["http_method"] = edit.HttpMethod
		params["tag"] = edit.Tag
		menu := model.NewPermissions()
		menu.UpdateMenu(id, params, c)
	}
	return
}

//DeletePermission 删除权限信息
//@Tags 后端权限模块
//@Summary 删除权限信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permission/{id} [delete]
func (a *Permission) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	menu := model.NewPermissions()
	menu.DeleteMenu(id, c)
	return
}

//GetTree 权限层级信息
//@Tags 后端权限模块
//@Summary 获取权限层级信息
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []map[string]interface{} "{"code":200,"data":[]map[string]interface{}}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/permissionTree [get]
func (a *Permission) GetTree(c *gin.Context) {
	menu := model.NewPermissions()
	menu.GetPermissionTree(c)
	return
}
