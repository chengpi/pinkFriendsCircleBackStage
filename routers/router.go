package routers

import (
	"smallRoutine/loveta/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //数据库交互测试接口
    beego.Router("/insertUserInfo",&controllers.TradeCreateController{},
    "post:InsertUserInfo")
    beego.Router("/readUserInfo",&controllers.TradeCreateController{},
    "get:ReadUserInfo")
    //预下单
	beego.Router("/trade_create",&controllers.TradeCreateController{},
	"post:Post")
}
