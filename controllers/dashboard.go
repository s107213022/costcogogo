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
	// 收新的 Flash 消息
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		// 將 Flash 消息傳遞給模板
		c.Data["FlashMessage"] = n
	}

	c.Data["Title"] = "Dashboard - CostcoGoGo"
	// 如果session中没有userID，跳回登入
	if userID == nil {
		c.Redirect("/", 302)
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
	for _, ow := range owelist {
		// 将 money 的值修改为 money * unit
		ow.Money = ow.Money * ow.Unit
	}
	// 計算owelist 的money*unit的總和
	owelistmoney := CalculateTotalAmount(owelist, 0)
	owelistmoneyovernotconfirm := CalculateTotalAmount(owelist, 1)
	fmt.Println(owelistmoney)
	paylist, err := models.GetOwelistByDebtorID(userID.(int64))
	if err != nil {
		panic(err)
	}
	for _, pw := range paylist {
		// 将 money 的值修改为 money * unit
		pw.Money = pw.Money * pw.Unit
	}
	paylistmoney := CalculateTotalAmount(paylist, 0)
	paylistmoneyovernotconfirm := CalculateTotalAmount(paylist, 1)

	c.Data["Owelist"] = owelist
	c.Data["Paylist"] = paylist
	c.Data["OwelistMoney"] = owelistmoney
	c.Data["OwelistMoneyOverNotConfirm"] = owelistmoneyovernotconfirm
	c.Data["PaylistMoney"] = paylistmoney
	c.Data["PaylistMoneyOverNotConfirm"] = paylistmoneyovernotconfirm
	c.Data["name"] = user.Name
	c.Data["account"] = user.Account
	c.TplName = "index.tpl"
}

func CalculateTotalAmount(oweList []*models.Owelist, status int64) int {
	totalAmount := 0
	for _, list := range oweList {
		// 计算 money * unit，并累加到 totalAmount
		if int64(list.Finish) == status {
			totalAmount += list.Money
		}
	}
	return totalAmount
}
