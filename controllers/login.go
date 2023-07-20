package controllers

import (
	"costcogogo/models"
	"fmt"

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
	account := c.GetString("account")
	password := c.GetString("password")
	loggedIn, userID, err := models.CheckLogin(account, password)
	if err != nil {
		panic(err)
	}
	if loggedIn {
		c.SetSession("userID", userID)
		c.Redirect("/dashboard", 302)
	}
	fmt.Println(loggedIn)
	fmt.Println(userID)
	c.Data["Title"] = "My Website"
	c.TplName = "login.tpl"
	// if account == "maple" && password == "123" {
	// 	c.Ctx.WriteString("登入成功")
	// } else {
	// 	c.Ctx.WriteString("用户名或密码错误")
	// }
}
