package controllers

import (
	"encoding/json"

	"github.com/alastairruhm/guidor/src/models"

	"net/http"

	"github.com/astaxie/beego"
)

// Operations about Users
type InstanceController struct {
	beego.Controller
}

func (i *InstanceController) Create() {
	var instance models.Instance
	json.Unmarshal(i.Ctx.Input.RequestBody, &instance)
	instance, err := models.CreateInstance(instance)
	if err != nil {
		i.CustomAbort(http.StatusBadRequest, "error create instance")
	}
	i.Data["json"] = &instance
	i.ServeJSON()
}

func (i *InstanceController) GetAll() {

}
