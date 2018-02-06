package controllers

import (
	"github.com/astaxie/beego"
	wxClass "wxpay/wxLib"
	"strconv"
	"strings"
	"time"
)


type PayController struct  {
	beego.Controller
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func (this PayController) Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}
	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

func (this PayController) Print(strMsg string) {
	this.Ctx.WriteString(strMsg)
	this.StopRun()
}

func (this PayController) GetTimestamp() int64{
	return time.Now().UnixNano() / 1000000 //毫秒
}

func (this PayController) GetClientIp() string {
	ip := this.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = this.Ctx.Request.RemoteAddr
	}
	if strings.Contains(ip, ":") {
		ip = this.Substr(ip, 0, strings.Index(ip, ":"))
	}
	return ip
}

func (this *PayController) Get() {
	//定义的三种方式
	//wx := &wxClass.WxController{}
	//wx := new(wxClass.WxController)
	wx := wxClass.WxController{}
	wx.RegisterWx()

	//获取openid
	var openid string

	wxCode := this.GetString("code")
	if wxCode == "" {
		this.Redirect(wx.CreateOauthUrlForCode("http://test3.beecloud.cn/wxpay/pay"), 302)
	}else{
		wx.SetCode(wxCode)

		var openidErr error
		openid, openidErr = wx.GetOpenid()
		if openidErr != nil {
			this.Print(openidErr.Error())
		}
	}

	timestamp := this.GetTimestamp()
	bill_no := "godemo" + strconv.FormatInt(timestamp, 10)

	wx.SetParameter("openid", openid)
	wx.SetParameter("out_trade_no", bill_no)
	wx.SetParameter("total_fee", "1")
	wx.SetParameter("trade_type", "JSAPI")
	wx.SetParameter("body", "微信支付测试demo")
	wx.SetParameter("notify_url", "http://www.example.cn/notify")
	wx.SetParameter("spbill_create_ip", this.GetClientIp())

	prepay_id, err := wx.GetPrepayId()
	if err != nil {
		this.Print(err.Error())
	}
	wx.SetPrepayId(prepay_id)

	this.Data["jsapi"] = wx.GetJsapiParameters()
	this.Data["channel"] = "JSAPI"
	this.TplName = "pay.tpl"
}

