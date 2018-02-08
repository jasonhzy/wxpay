package routers

import (
	"github.com/astaxie/beego"
	"wxpay/controllers"
	_"wxpay/init"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/wxpay/demo", &controllers.PayController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/session", &controllers.UserController{}, "*:Session")
	beego.Router("/user/add", &controllers.UserController{}, "*:Add")
	beego.Router("/user/update", &controllers.UserController{}, "*:Update")
}
