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
	c.TplName = "signin.tpl"
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

		// 檢查是否勾選 Remember password
		rememberPassword := c.GetString("rememberPasswordCheck") == "on"
		if rememberPassword {
			// 如果勾選了 Remember password，將帳號資訊存入 Cookie 中，有效期 30 天
			c.Ctx.SetCookie("account", account, 3600*24*30)
		} else {
			// 如果沒有勾選 Remember password，則刪除 Cookie 中的帳號資訊
			c.Ctx.SetCookie("account", "", -1)
		}

		c.Redirect("/dashboard", 302)
	}

	fmt.Println(loggedIn)
	fmt.Println(userID)
	c.Data["Title"] = "My Website"
	c.TplName = "login.tpl"
}
