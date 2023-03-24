package web

import (
	"fmt"
	"gin-icqqg/model"
	"github.com/gin-gonic/gin"
)

type AddMenu struct {
	ParentID int64  `form:"parent_id" json:"parent_id"`            // 父级ID
	Sort     int64  `form:"sort" json:"sort" binding:"required"`   //  菜单排序
	Title    string `form:"title" json:"title" binding:"required"` // 菜单名称
	Url      string `form:"url" json:"url" binding:"required"`     // 路由地址
	Show     int64  `form:"show" json:"show" binding:"required"`   // 是否展示 1是 2否
}
type EditMenu struct {
	ParentID int64  `form:"parent_id" json:"parent_id"` // 父级ID
	Sort     int64  `form:"sort" json:"sort" `          //  菜单排序
	Title    string `form:"title" json:"title" `        // 菜单名称
	Url      string `form:"url" json:"url" `            // 路由地址
	Show     int64  `form:"show" json:"show"`           // 是否展示 1是 2否
}

func NewAddMenu() *AddMenu {

	return &AddMenu{}
}

//List 数据列表
//@Tags 后端菜单模块
//@Summary 菜单列表
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []model.Menu "{"code":200,"data":[]model.Menu}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menuList [get]
func (a *AddMenu) List(c *gin.Context) {
	menu := model.NewMenu()
	menu.List(c)
	return
}

//Get 通过ID获取菜单信息
//@Tags 后端菜单模块
//@Summary 获取菜单信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} model.Menu "{"code":200,"data":model.Menu}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menu/{id} [get]
func (a *AddMenu) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
	} else {
		menu := model.NewMenu()
		menu.GetOne(id, c)
	}
	return
}

//Add 新增菜单信息
//@Tags 后端菜单模块
//@Summary 添加菜单信息
//@Param token header string true "token"
//@Param sort formData int true "排序升序"
//@Param title formData string true "菜单标题"
//@Param url formData string true "菜单路由"
//@Param show formData int true "是否显示 1是 2否"
//@Param parent_id formData int true "父级ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menu [post]
func (a *AddMenu) Add(c *gin.Context) {
	params := map[string]interface{}{}
	err := c.ShouldBind(&a)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["show"] = a.Show
		params["parent_id"] = a.ParentID
		params["sort"] = a.Sort
		params["title"] = a.Title
		params["url"] = a.Url
		menu := model.NewMenu()
		menu.AddMenu(params, c)
	}
	return
}

//Edit 修改菜单信息
//@Tags 后端菜单模块
//@Summary 修改菜单信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Param sort formData int false "排序升序"
//@Param title formData string false "菜单标题"
//@Param url formData string false "菜单路由"
//@Param show formData int false "是否显示 1是 2否"
//@Param parent_id formData int false "父级ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menu/{id} [put]
func (a *AddMenu) Edit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	params := map[string]interface{}{}
	var edit EditMenu
	err := c.ShouldBind(&edit)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("%v", err)})
	} else {
		params["show"] = edit.Show
		params["parent_id"] = edit.ParentID
		params["sort"] = edit.Sort
		params["title"] = edit.Title
		params["url"] = edit.Url
		menu := model.NewMenu()
		menu.UpdateMenu(id, params, c)
	}
	return
}

//Delete 删除菜单信息
//@Tags 后端菜单模块
//@Summary 删除菜单信息
//@Param token header string true "token"
//@Param id path int true "ID"
//@Produce json
// @Success 200 {object} response.Code "{"code":200,"msg":"success"}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menu/{id} [delete]
func (a *AddMenu) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "参数有误"})
		c.Abort()
		return
	}
	menu := model.NewMenu()
	menu.DeleteMenu(id, c)
	return
}

//GetTree 菜单层级信息
//@Tags 后端菜单模块
//@Summary 获取菜单层级信息
//@Param token header string true "token"
//@Produce json
// @Success 200 {object} []map[string]interface{} "{"code":200,"data":[]map[string]interface{}}"
// @Failure 400 {object} response.Code "请求错误"
// @Failure 500 {object} response.Code "内部错误"
//@Router /api/web/menuTree [get]
func (a *AddMenu) GetTree(c *gin.Context) {
	menu := model.NewMenu()
	menu.GetMenuTree(c)
	return
}
