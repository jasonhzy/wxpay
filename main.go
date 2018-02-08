package main

import (
	_ "wxpay/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	//Debug 为 true 打印查询的语句, 不建议使用在生产模式
	orm.Debug = true

	//设置是否开启 Session，默认是 false，配置文件对应的参数名：sessionon。
	//beego.BConfig.WebConfig.Session.SessionOn = true

	//设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID，配置文件对应的参数名是：sessionname。
	//beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"

	//设置 Session 过期的时间，默认值是 3600 秒，配置文件对应的参数：sessiongcmaxlifetime。
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 10

	//设置cookie的过期时间，cookie是用来存储保存在客户端的数据。默认值是 0，即浏览器生命周期
	//beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 0

	beego.Run()
}

