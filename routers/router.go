package routers

import (
	"costcogogo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
}
