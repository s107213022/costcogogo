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

	// id取name

	name, err := models.GetNameById(userID.(int64))
	if err != nil {
		panic(err)
	}
	owelist, err := models.GetOwelistByCreditorID(userID.(int64))
	if err != nil {
		panic(err)
	}
	// 遍历 owelist 切片，替换 Creditor 和 Debtor 字段为 CreditorName 和 DebtorName
	// for _, ow := range owelist {
	// 	// name, err := models.GetNameById(ow.Creditor)
	// 	fmt.Println(ow.Debtor)
	// }

	// c.Data["Owelist"] = owelist
	// fmt.Println(owelist)
	for _, ow := range owelist {
		fmt.Println("ID:", ow.Id)
		fmt.Println("Creditor ID:", ow.Creditor.Id)
		fmt.Println("Creditor Name:", ow.Creditor.Name)
		fmt.Println("Items:", ow.Items)
		fmt.Println("Money:", ow.Money)
		fmt.Println("Unit:", ow.Unit)
		fmt.Println("Date:", ow.Date)
		fmt.Println("Finish:", ow.Finish)
		fmt.Println("Debtor ID:", ow.Debtor.Id)
		fmt.Println("Debtor Name:", ow.Debtor.Name)
		fmt.Println("-------------")
	}
	// c.Data["Owelist"] = list
	// All ,err := models.GetAllOwelist()
	c.Data["name"] = name
	c.TplName = "dashboard.tpl"
}
