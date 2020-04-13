package controllers

import (
	"github.com/beego/i18n"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}
func (this *BaseController) Prepare() {
	lang := this.Ctx.Input.Header("language")
	if lang == "en" {
		this.Lang = "en-US"
	} else if lang == "tw" {
		this.Lang = "zh-TW"
	} else {
		this.Lang = "zh-CN"
	}
}