package controllers

import (
	"costcogogo/models"

	"github.com/astaxie/beego"
)

// DashboardController operations for Dashboard
type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {
	userID := c.GetSession("userID")
	c.Data["Title"] = "My Website"
	// 如果session中没有userID，跳回登入
	if userID == nil {
		c.Redirect("/login", 302)
		return
	}
	// id取name

	name, err := models.GetNameById(userID.(int64))
	if err != nil {
		panic(err)
	}
	c.SetSession("name", name)
	owelist, err := models.GetOwelistByCreditorID(userID.(int64))
	if err != nil {
		panic(err)
	}
	c.Data["Owelist"] = owelist
	// All ,err := models.GetAllOwelist()
	c.Data["name"] = name
	c.TplName = "dashboard.tpl"
}
