package controllers
import (
"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {

	c.TplName = "index.tpl"
}

