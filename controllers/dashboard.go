package controllers

import (
	"github.com/astaxie/beego"
)

// DashboardController operations for Dashboard
type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	userID := c.GetSession("userID")
	c.Data["Title"] = "My Website"
	// All ,err := models.GetAllOwelist()
	c.Data["name"] = userID
	c.TplName = "dashboard.tpl"
}
