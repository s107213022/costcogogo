package controllers

import (
	"github.com/astaxie/beego"
)

// OwemoneyController operations for Owemoney
type OwemoneyController struct {
	beego.Controller
}

// URLMapping ...
func (c *OwemoneyController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Owemoney
// @Param	body		body 	models.Owemoney	true		"body for Owemoney content"
// @Success 201 {object} models.Owemoney
// @Failure 403 body is empty
// @router / [post]
func (c *OwemoneyController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Owemoney by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Owemoney
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OwemoneyController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Owemoney
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Owemoney
// @Failure 403
// @router / [get]
func (c *OwemoneyController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Owemoney
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Owemoney	true		"body for Owemoney content"
// @Success 200 {object} models.Owemoney
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OwemoneyController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Owemoney
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OwemoneyController) Delete() {

}
