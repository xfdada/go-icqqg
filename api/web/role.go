package web

import (
	"fmt"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type Role struct {
	RoleName      string `form:"role_name" gorm:"column:role_name;type:varchar(50);" json:"role_name,omitempty" binding:"required"` // 角色名称
	RoleTag       string `form:"role_tag" gorm:"column:role_tag;type:varchar(50);" json:"role_tag,omitempty" binding:"required"`    // 角色标签
	MenuID        string `form:"menu_id" gorm:"column:menu_id;type:varchar(255);" json:"menu_id,omitempty"`                         // 菜单ID    用逗号分隔例如 1，2，3，4
	PermissionsID string `form:"permissions_id" gorm:"column:permissions_id;type:varchar(255);" json:"permissions_id,omitempty"`    // 权限ID    用逗号分隔例如 1，2，3，4
}

type EditRole struct {
	RoleName      string `form:"role_name" gorm:"column:role_name;type:varchar(50);" json:"role_name,omitempty"`                 // 角色名称
	RoleTag       string `form:"role_tag" gorm:"column:role_tag;type:varchar(50);" json:"role_tag,omitempty"`                    // 角色标签
	MenuID        string `form:"menu_id" gorm:"column:menu_id;type:varchar(255);" json:"menu_id,omitempty"`                      // 菜单ID    用逗号分隔例如 1，2，3，4
	PermissionsID string `form:"permissions_id" gorm:"column:permissions_id;type:varchar(255);" json:"permissions_id,omitempty"` // 权限ID    用逗号分隔例如 1，2，3，4
}

func NewRole() *Role {
	return &Role{}
}

//List 数据列表
//@Tags 后端角色模块
//@Summary 角色列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.Role "{"code":200,"data":[]model.Menu}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/roleList [get]
func (a *Role) List(c *gin.Context) {
	menu := model.NewRole()
	menu.List(c)
	return
}

//Get 通过ID获取角色信息
//@Tags 后端角色模块
//@Summary 获取角色信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} model.Role "{"code":200,"data":model.Role}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/role/{id} [get]
func (a *Role) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		menu := model.NewRole()
		menu.GetOne(id, c)
	}
	return
}

//Add 新增角色信息
//@Tags 后端角色模块
//@Summary 添加角色信息
//@Param token header string true "token"
//@Param role_name formData int true "角色英文名"
//@Param role_tag formData string true "角色中文名"
//@Param menu_id formData string true "菜单ID字符串1，2，3"
//@Param permissions_id formData string true "权限ID字符串 1，2，3"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/role [post]
func (a *Role) Add(c *gin.Context) {
	params := map[string]interface{}{}
	err := c.ShouldBind(&a)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["name"] = a.RoleName
		params["tag"] = a.RoleTag
		params["menu_id"] = a.MenuID
		params["permissions_id"] = a.PermissionsID
		menu := model.NewRole()
		menu.AddRole(params, c)
	}
	return
}

//Edit 修改菜单信息
//@Tags 后端菜单模块
//@Summary 修改菜单信息
//@Param token header string true "token"
//@Param role_name formData int true "角色英文名"
//@Param role_tag formData string true "角色中文名"
//@Param menu_id formData string true "菜单ID字符串1，2，3"
//@Param permissions_id formData string true "权限ID字符串 1，2，3"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/role/{id} [put]
func (a *Role) Edit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	params := map[string]interface{}{}
	var edit EditRole
	err := c.ShouldBind(&edit)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["name"] = edit.RoleName
		params["tag"] = edit.RoleTag
		params["menu_id"] = edit.MenuID
		params["permissions_id"] = edit.PermissionsID
		menu := model.NewRole()
		menu.UpdateRole(id, params, c)
	}
	return
}

//Delete 删除角色信息
//@Tags 后端角色模块
//@Summary 删除角色信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/role/{id} [delete]
func (a *Role) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	menu := model.NewRole()
	menu.DeleteRole(id, c)
	return
}
