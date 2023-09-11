package routers

import (
	"costcogogo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/create", &controllers.CreateAccountController{})
	beego.Router("/newowe", &controllers.NewoweController{})
	beego.Router("/updateowe/:id/", &controllers.UpdateowelistController{})
	beego.Router("/updateowe/:id/changestatus", &controllers.UpdateoweController{}, "post:ChangeStatus")
	beego.Router("/updateowe/:id/payover", &controllers.UpdateoweController{}, "post:Payover")
	beego.Router("/updateowe/:id/delete", &controllers.UpdateoweController{}, "post:Delete")
}
