package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
	err  error
	data interface{}
}

// 函数结束时,组装成json结果返回
func (c *APIController) Finish() {
	r := struct {
		Error interface{} `json:"error"`
		Data  interface{} `json:"data"`
	}{}
	if c.err != nil {
		r.Error = c.err.Error()
	}
	r.Data = c.data
	c.Data["json"] = r
	c.ServeJSON()
}

// 如果请求的参数不存在,就直接 error返回
func (c *APIController) MustString(key string) string {
	v := c.GetString(key)
	if v == "" {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("require filed: %s", key),
			"data":  "orz!!",
		}
		c.ServeJSON()
		c.StopRun()
	}
	return v
}

// TODO
