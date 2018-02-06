package routers

import (
	"wxpay/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/wxpay/demo", &controllers.PayController{})
}
