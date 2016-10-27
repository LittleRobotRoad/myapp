package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController  struct {
	beego.Controller
	i18n.Locale
}

type MainController struct {
	baseController
}

func (this *baseController)Prepare() {
	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

func (this *MainController)Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "beego.me"
	this.Data["Hi"] = "hi"
	this.Data["Bye"] = "bye"
	this.TplName = "index.html"
}