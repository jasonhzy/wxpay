package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
)

type CommonController struct {
	beego.Controller
}

func (this CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
	return
}

func (this CommonController) AjaxRsp(data map[string]interface{}, count int64) {
	this.Data["json"] = &map[string]interface{}{"data": data, "count": count}
	this.ServeJSON()
	return
}

func (this CommonController) GetMd5(str string) string {
	sign := md5.New()
	sign.Write([]byte(str))
	return hex.EncodeToString(sign.Sum(nil))
}