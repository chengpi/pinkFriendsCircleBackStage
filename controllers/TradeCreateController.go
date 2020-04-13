package controllers

import (
	"fmt"
	"smallRoutine/loveta/models/tt_pay/trade"
	"time"
	"math/rand"
	"strconv"
	"context"
	//"encoding/json"
	//"net/url"
	"encoding/json"
	"smallRoutine/loveta/models"
	//"math"
)

type TradeCreateController struct {
	BaseController
}

//func (this *TradeCreateController)Get() {
//
//}
//func (this *TradeCreateController)Post() {
//	username := this.GetString("username")
//	password := this.GetString("password")
//	fmt.Println("username:", username, ",password:", password)
//	this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功","status":"success"}
//	this.ServeJSON()
//}
func (this *TradeCreateController)InsertUserInfo() {
	var userInfo models.UserInfo
	//userInfo.Id,_ = strconv.ParseInt(this.GetString("id"),10,64)
	//userInfo.Id,_ = this.GetInt64("id")
	userInfo.UserName = this.GetString("user_name")
	//userInfo.TotalRewards,_ = this.GetFloat("total_rewards",0.00)

	totalRewards,_ := this.GetFloat("total_rewards")
	//fmt.Println(totalRewards)
	//fmt.Println(math.Round(totalRewards*100)/100)
	userInfo.TotalRewards,_ = strconv.ParseFloat(fmt.Sprintf("%.2f", totalRewards), 64)
	//fmt.Println(fmt.Sprintf("%.2f", totalRewards))
	//fmt.Println(userInfo.TotalRewards)

	settledAmounts,_ := this.GetFloat("settled_amounts")
	//fmt.Println(settledAmounts)
	//fmt.Println(math.Round(settledAmounts*100)/100)
	userInfo.SettledAmounts,_ = strconv.ParseFloat(fmt.Sprintf("%.2f", settledAmounts), 64)
	//fmt.Println(fmt.Sprintf("%.2f", settledAmounts))
	//fmt.Println(userInfo.SettledAmounts)

	userInfo.HeadLogo = this.GetString("head_logo")
	userInfo.CustomerId = this.GetString("customer_id")
	userInfo.OpenId = this.GetString("open_id")


	str,err := models.InsertOrUpdateUserInfo0(userInfo)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": fmt.Sprintf("%v",err),
			"status":"false",
		}

		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": 1,
		"message": str,
		"status":"true",
	}

	this.ServeJSON()
	return


}
func (this *TradeCreateController)ReadUserInfo()  {
	openId := this.GetString("open_id")
	userInfo, err :=models.QueryUserInfo(openId)
	if err != nil{
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": fmt.Sprintf("%v",err),
			"status":"false",
		}

		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{
		"code": 1,
		"message": userInfo,
		"status":"true",
	}

	this.ServeJSON()
	return

}

func (this *TradeCreateController)Post() {
	//username := this.GetString("username")
	//password := this.GetString("password")
	//fmt.Println("username:", username, ",password:", password)
	//this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功","status":"success"}
	rand.Seed(time.Now().UnixNano())
	totalAmount,_ := strconv.Atoi(this.GetString("total_amount"))

	req := trade.NewTradeCreateRequest(conf)

	// 下面两个版本号，AppletVersion指的是小程序收银台版本，Version指的是财经后端下单接口版本
	// 小程序收银台版本有1.0和2.0，头条APP只在7.2.7之后支持收银台2.0版本，7.2.7之前的版本请使用1.0
	// 该参数可设置为"1.0"(返回拉起1.0收银台参数)，"2.0"(返回拉起2.0收银台参数),"2.0+"(返回一个json，包含1.0和2.0参数)
	// 小程序收银台版本，可选1.0，2.0及2.0+
	req.AppletVersion = "2.0"
	// 后端下单接口默认为2.0， 可更改为1.0
	req.Version = "2.0"
	// 此处是随机生成的，使用时请填写您的商户订单号
	req.OutOrderNo = fmt.Sprintf("%d%d", time.Now().Unix(),rand.Intn(100000000))
	// 填写用户在头条的id
	req.Uid = this.GetString("uid")
	// 填写订单金额
	req.TotalAmount = totalAmount
	// 填写币种，一般均为CNY
	req.Currency = "CNY"
	// 固定值，不要改动
	req.TradeType = "H5"
	// 填写您的订单名称
	req.Subject = this.GetString("subject")
	// 填写您的订单内容
	req.Body = this.GetString("body")
	// 交易时间，此处自动生成，您也可以根据需求赋值，但必须为Unix时间戳
	req.TradeTime = fmt.Sprintf("%d", time.Now().Unix())
	// 填写您的订单有效时间（单位：秒）
	req.ValidTime = "36000"
	// 填写您的异步通知地址
	req.NotifyUrl = "https://google.com"
	// 严格json字符串格式
	req.RiskInfo = `{"ip":"127.0.0.1", "device_id":"122333"}`
	// 固定值，不要改动
	req.ProductCode = "pay"

	// 1.0版本特有参数，当AppletVersion填"1.0"和"2.0+"时需要填写
	req.Params = `{"url":"..."}`      // 传递给支付方的支付信息，标准 json 格式字符串，不同的支付方参数格式不一样
	req.PaymentType1_0 = "ALIPAY_APP" // 1.0版本的PaymentType，目前只支持支付宝，请填写ALIPAY_APP
	req.PayChannel = "ALIPAY_NO_SIGN" // 目前只支持支付宝，请填写ALIPAY_NO_SIGN

	// 2.0版本特有参数，当AppletVersion填"2.0"和"2.0+"时需要填写
	req.PaymentType = "direct" // 2.0版本的PaymentType，固定值direct，不要改动，
	req.WxUrl = "https://wx.tenpay.com/cgi-bin/mmpayweb-bin/checkmweb"
	// 调用微信H5支付统一下单接口返回的mweb_url字段值。service=1时(外部开发者)必传，否则无法使用微信支付
	req.WxType = "MWEB"        // service=1时(外部开发者)，且wx_url有值时，传固定值：MWEB
	req.AlipayUrl = "alipay_sdk"
	// 调用支付宝App支付的签名订单信息，详见支付宝App支付请求参数说明。
	// service=1时(外部开发者)必传，否则无法使用支付宝支付

	//fmt.Println("total_amount:", totalAmount, ",uid:", req.Uid)
	//fmt.Println("subject:", req.Subject, ",body:", req.Body)
	//fmt.Println("app_id:",req.Config.AppId)
	ctx := context.Background()
	resp, err := trade.TradeCreate(ctx, req)
	//fmt.Println("***1***")
	//valueErr := url.Values{}
	//json.Unmarshal([]byte(fmt.Sprintf("%v",err)),valueErr)
	//errBytes,_ := json.Marshal(err)
	//fmt.Println(string(errBytes))

	if err != nil {
		//c.String(http.StatusOK, fmt.Sprintf("Request failed! \nRequest:\n  [%#v]\nError:\n  [%s]\n", req, err))
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"message": fmt.Sprintf("%v",err),
			"status":"fail",
		}
		//fmt.Println("***20***")
		//fmt.Println(err)

		this.ServeJSON()
		return
	}
	//fmt.Println("***21***")
	cashDeskParams, _ := resp.GetCashdeskAppletParams()
	//values := url.Values{}
	//err = json.Unmarshal([]byte(cashDeskParams),&values)
	//fmt.Println(values)
	//if err != nil {
	//	this.Data["json"] = map[string]interface{}{
	//		"code": 0,
	//		"message": fmt.Sprintf("%v",err),
	//		"status":"fail",
	//	}
	//	this.ServeJSON()
	//	return
	//}
	params := make(map[string]string)
	values := make(map[string]string)
	outputParams := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(cashDeskParams),&params)//注意要有取址符
	for i,v := range params{
		err = json.Unmarshal([]byte(v),&values)
		outputParams[i] = values
	}

	this.Data["json"] = map[string]interface{}{
		"code": 1,
		"message": outputParams,
		"status":"success",
	}
	this.ServeJSON()
}

//报错：Handler crashed with error can't find templatefile
// in the path:views/tradecreatecontroller/post.tpl
//每有一行this.Data["json"] = map[string]interface{}{"code": 0, "message": err,"status":"fail"}
//就要配这么一行this.ServeJSON()
//不然就报错这个
