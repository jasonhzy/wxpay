package controllers

import (
	m "wxpay/models"
	"fmt"
	"time"
)

type UserController struct{
	CommonController
}

func (this *UserController) Get() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "create_time"
	}
	users, count := m.GetList(page, page_size, sort)
	//beego 中使用 session 相当方便，只要在 main 入口函数中设置如下：
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//或者通过配置文件配置如下：
	//sessionon = true
	this.SetSession("users", users)
	if this.IsAjax() {
		this.AjaxRsp(map[string]interface{}{"total": count, "rows": users}, count)
	} else {
		this.Data["data"] = users
		this.Data["count"] = count
		this.TplName = "user.tpl"
	}
}

func (this *UserController) Session() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+ "%s\n", this.GetSession("users"))
	//销毁session
	//this.DelSession("user")
	//this.DestroySession()
	this.StopRun()
}

func (this *UserController) Add() {
	u := m.User{}
	//if err := this.ParseForm(&u); err != nil {
	//	//handle error
	//	this.Rsp(false, err.Error())
	//	return
	//}
	u.Username = "jason"
	u.Password = this.GetMd5("123456")
	u.Nickname = "jason hu"
	u.Age = 20

	id, err := m.Add(u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
	} else {
		this.Rsp(false, err.Error())
	}
}

func (this *UserController) Update() {
	u := m.User{}
	u.Id = 2

	err := m.GetUserById(u)
	if err != nil {
		this.Rsp(false, err.Error())
	}

	u.Age = 20
	num, err := m.UpdateBycolumn(u, "Age")
	if err == nil && num > 0 {
		this.Rsp(true, "Success")
	} else {
		this.Rsp(false, err.Error())
	}
}

func (this *UserController) Del() {
	u := m.User{}
	u.Id = 2

	err := m.GetUserById(u)
	if err != nil {
		this.Rsp(false, err.Error())
	}

	u.Age = 20
	num, err := m.UpdateBycolumn(u, "Age")
	if err == nil && num > 0 {
		this.Rsp(true, "Success")
	} else {
		this.Rsp(false, err.Error())
	}
}