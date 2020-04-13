package controllers

import (
	"smallRoutine/loveta/models/tt_pay/config"
	"github.com/astaxie/beego"
)

var conf config.Config
func init()  {
	 conf = config.Config{
		AppId:             beego.AppConfig.String("AppId"),
		// 支付方分配给业务方的ID，用于获取 签名/验签 的密钥信息
		AppSecret:         beego.AppConfig.String("AppSecret"),
		// 支付方密钥
		MerchantId:        beego.AppConfig.String("MerchantId"),
		// 支付方分配给业务方的商户编号
		TPDomain:          beego.AppConfig.String("TPDomain"),

		TPClientTimeoutMs: 6000,
	}
	
}
