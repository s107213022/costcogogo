package controllers

import (
	"costcogogo/models"
	"fmt"

	"github.com/astaxie/beego"
)

// CreateAccountController operations for CreateAccount
type CreateAccountController struct {
	beego.Controller
}

func (c *CreateAccountController) Get() {
	c.Data["Title"] = "Create Your Account"
	c.TplName = "createAccount.tpl"
}

func (c *CreateAccountController) Post() {
	name := c.GetString("name")
	account := c.GetString("account")
	password := c.GetString("password")
	userAdd := models.User{Name: name, Account: account, Password: password}
	userid, err := models.AddUser(&userAdd)
	if err != nil {
		panic(err)
	}
	fmt.Println(userid)
	c.Redirect("/login", 302)
}
