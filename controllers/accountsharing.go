package controllers

import (
	"costcogogo/models"

	"github.com/astaxie/beego"
)

// CreateAccountController operations for CreateAccount
type AccountsharingController struct {
	beego.Controller
}

func (c *AccountsharingController) Get() {
	c.Data["Title"] = "Create Your Accountsharing"
	c.TplName = "accountsharing-form-elements.tpl"
	userID := c.GetSession("userID")
	if userID == nil {
		c.Redirect("/", 302)
		return
	}
	// user 是一串
	user, err := models.GetUserById(userID.(int64))
	if err != nil {
		// 處理錯誤
		panic(err)
	}
	NotmeUsers, err := models.GetNamesExceptSelf(userID.(int64))
	if err != nil {
		// 處理錯誤
		panic(err)
	}

	c.Data["Users"] = NotmeUsers
	c.Data["name"] = user.Name
	c.Data["account"] = user.Account
}

func (c *AccountsharingController) Post() {

}
