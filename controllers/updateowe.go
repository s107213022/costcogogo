package controllers

import (
	"costcogogo/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// CreateAccountController operations for CreateAccount
type UpdateoweController struct {
	beego.Controller
}



func (c *UpdateoweController) ChangeStatus() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	owelist, err := models.GetOwelistById(id)
	if err != nil {
		c.Data["json"] = map[string]string{"status": "error", "message": "Owelist not found"}
	} else {
		// 將 Finish 狀態從 1 修改為 2
		owelist.Finish = 2
		err := models.UpdateOwelist(owelist)
		if err != nil {
			c.Data["json"] = map[string]string{"status": "error", "message": "Failed to update Owelist status"}
		} else {
			c.Data["json"] = map[string]string{"status": "success", "message": "Owelist status updated successfully"}
		}
	}

	c.ServeJSON()
}

func (c *UpdateoweController) Payover() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	owelist, err := models.GetOwelistById(id)
	if err != nil {
		c.Data["json"] = map[string]string{"status": "error", "message": "Owelist not found"}
	} else {
		// 將 Finish 狀態從 0 修改為 1
		owelist.Finish = 1
		err := models.UpdateOwelist(owelist)
		if err != nil {
			c.Data["json"] = map[string]string{"status": "error", "message": "Failed to update Owelist status"}
		} else {
			c.Data["json"] = map[string]string{"status": "success", "message": "Owelist status updated successfully"}
		}
	}

	c.ServeJSON()
}

func (c *UpdateoweController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	fmt.Println(id)
	err := models.DeleteOwelist(id)
	if err != nil {
		// 處理錯誤，如：顯示錯誤消息或將錯誤記錄到log中
		c.Data["json"] = map[string]string{"status": "error", "message": "Failed to delete Owelist"}
	} else {
		c.Data["json"] = map[string]string{"status": "success", "message": "Owelist deleted successfully"}
	}

	c.ServeJSON()
}
