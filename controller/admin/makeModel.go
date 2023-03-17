package admin

import (
	"bytes"
	"fmt"
	"gin-icqqg/config"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func Helpers(c *gin.Context) {

	Model := "Product"
	var Filed []map[string]interface{}
	list1 := map[string]interface{}{"Name": "Name", "Type": "string", "Column": "name", "SqlType": "varchar", "Size": "50", "Comment": "名称"}
	list2 := map[string]interface{}{"Name": "Title", "Type": "string", "Column": "title", "SqlType": "varchar", "Size": "50", "Comment": "标题"}
	list3 := map[string]interface{}{"Name": "Version", "Type": "string", "Column": "version", "SqlType": "varchar", "Size": "50", "Comment": "型号"}
	Filed = append(Filed, list1, list2, list3)
	Data := map[string]interface{}{
		"Model":     Model,
		"Filed":     Filed,
		"TableName": "table_product",
		"Name":      "产品",
		"Title":     "后台产品",
		"Path":      "product",
	}

	tempModel, _ := template.New("createModel.tmpl").ParseFiles("resource/view/admin/createModel.tmpl")
	tempService, _ := template.New("createService.tmpl").ParseFiles("resource/view/admin/createService.tmpl")

	var ModelBytes bytes.Buffer
	var ServiceBytes bytes.Buffer
	if err := tempModel.Execute(&ModelBytes, Data); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error executing template: %v", err))
		return
	}
	if err := tempService.Execute(&ServiceBytes, Data); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error executing template: %v", err))
		return
	}
	err := ioutil.WriteFile("model/"+strings.ToLower(Model)+".go", ModelBytes.Bytes(), 0777)
	err = ioutil.WriteFile("api/web/"+strings.ToLower(Model)+".go", ServiceBytes.Bytes(), 0777)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("%v", err))
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "模板输出成功"})

}
