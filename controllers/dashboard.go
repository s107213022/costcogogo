package controllers

import (
	"costcogogo/models"
	"fmt"

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

	user, err := models.GetUserById(userID.(int64))
	if err != nil {
		// 處理錯誤
		panic(err)
	}
	fmt.Println(user.Name)
	c.SetSession("name", user.Name)
	owelist, err := models.GetOwelistByCreditorID(userID.(int64))
	if err != nil {
		panic(err)
	}
	c.Data["Owelist"] = owelist
	// All ,err := models.GetAllOwelist()
	c.Data["name"] = user.Name
	c.Data["account"] = user.Account
	c.TplName = "index.tpl"
}
