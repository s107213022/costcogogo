package controllers

import (
	"costcogogo/models"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// CreateAccountController operations for CreateAccount
type NewoweController struct {
	beego.Controller
}

func (c *NewoweController) Get() {
	c.Data["Title"] = "Create Your owe"
	c.TplName = "newowe-form-elements.tpl"
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

func (c *NewoweController) Post() {
	debtorstr := c.GetString("debtor")
	items := c.GetString("items")
	money := c.GetString("money")
	unit := c.GetString("unit")
	userID := c.GetSession("userID")
	var v models.Owelist
	moneyInt, err := strconv.Atoi(money)
	if err != nil {
		panic(err)
	}
	unitInt, err := strconv.Atoi(unit)
	if err != nil {
		panic(err)
	}

	Creditor, err := models.GetUserById(userID.(int64))
	if err != nil {
		panic(err)
	}
	debtor, err := models.GetUserByName(debtorstr)
	if err != nil {
		panic(err)
	}
	v.Creditor = Creditor
	v.Items = items
	v.Money = moneyInt
	v.Unit = unitInt
	v.Date = time.Now()
	v.Finish = 0
	v.Debtor = debtor

	oweid, err := models.AddoweList(&v)

	if err != nil {
		panic(err)
	}
	fmt.Println(oweid)
	fmt.Println(Creditor.Name)
	//c.Data["name"] = Creditor.Name
	// 新flash訊息
	flash := beego.NewFlash()
	flash.Notice("欠款單新增成功")
	flash.Store(&c.Controller) // 儲存 Flash 消息

	c.Redirect("/dashboard", 302)
}
