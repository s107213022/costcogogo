package routers

import (
	"costcogogo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/create", &controllers.CreateAccountController{})
	beego.Router("/newowe", &controllers.NewoweController{})
	beego.Router("/updateowe/:id/", &controllers.UpdateoweController{})
}
