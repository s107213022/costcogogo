package controllers

import (
	"github.com/astaxie/beego"
)

// DashboardController operations for Dashboard
type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	username := c.GetSession("username")
	c.Data["Title"] = "My Website"
	c.Data["name"] = username
	c.TplName = "dashboard.tpl"
}
