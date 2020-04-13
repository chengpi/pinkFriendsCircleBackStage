package main

import (
	_ "smallRoutine/loveta/routers"
	"github.com/astaxie/beego"
	_ "github.com/Go-SQL-Driver/MySQL"//使用orm包来和数据库做交互时记得引入这个驱动包
	//"smallRoutine/loveta/models/tt_pay/config"
	"github.com/astaxie/beego/orm"
)
//var conf config.Config

//func init()  {
//	 conf = config.Config{
//		AppId:             beego.AppConfig.String("AppId"),
//		// 支付方分配给业务方的ID，用于获取 签名/验签 的密钥信息
//		AppSecret:         beego.AppConfig.String("AppSecret"),
//		// 支付方密钥
//		MerchantId:        beego.AppConfig.String("MerchantId"),
//		// 支付方分配给业务方的商户编号
//		TPDomain:          beego.AppConfig.String("TPDomain"),
//
//		TPClientTimeoutMs: 6000,
//	}
//}
func init()  {
	path := beego.AppConfig.String("mysqluser") + "" +
		":" + beego.AppConfig.String("mysqlpass") + "" +
			"@tcp(" + beego.AppConfig.String("mysqlhost") + "" +
				":" + beego.AppConfig.String("mysqlport") + ")/" +
					"" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置最大空闲连接
	maxIdle := 30
	//设置最大空闲连接
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", path, maxIdle, maxConn)

}
func main() {
	//beego.AutoRender = false
	beego.SetStaticPath("/static", "static")
	beego.Run()
}

