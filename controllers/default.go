package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "jasonhzy.github.io"
	c.Data["Email"] = "lotushzy@163.com"
	c.TplName = "index.tpl"
}
