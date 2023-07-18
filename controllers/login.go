package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Title"] = "My Website"
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	name := c.GetString("name")
	password := c.GetString("password")
	if name == "maple" && password == "123" {
		c.Ctx.WriteString("登入成功")
	} else {
		c.Ctx.WriteString("用户名或密码错误")
	}
}
