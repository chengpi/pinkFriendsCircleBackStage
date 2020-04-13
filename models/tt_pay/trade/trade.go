package trade

import (
	"github.com/bitly/go-simplejson"
	"smallRoutine/loveta/models/tt_pay/config"
	"smallRoutine/loveta/models/tt_pay/consts"
	"fmt"
	"time"
	"net/url"
	"smallRoutine/loveta/models/tt_pay/util"
	"encoding/json"
)

//*********TradeCreateRequest*********//
// 预下单Request
type TradeCreateRequest struct {
	config.Config
	Method         string//*
	Format         string//*
	Charset        string//*
	SignType       string//*
	Timestamp      string//*
	Version        string//
	AppletVersion  string//
	bizContent     *simplejson.Json//* ///
	OutOrderNo     string//
	Uid            string//
	UidType        string
	TotalAmount    int   //
	Currency       string//
	TradeType      string//
	Subject        string//
	Body           string//
	ProductCode    string//
	TradeTime      string//
	ValidTime      string//
	NotifyUrl      string//
	RiskInfo       string//
	PaymentType    string
	PaymentType1_0 string
	Params         string
	ProductId      string
	PayChannel     string
	PayDiscount    string
	ServiceFee     string
	Path           string//*
	AlipayUrl      string
	WxUrl          string
	WxType         string
	ExtParam       string
}
func NewTradeCreateRequest(config config.Config) *TradeCreateRequest {
	ret := new(TradeCreateRequest)
	ret.Config = config
	ret.Method = consts.MethodTradeCreate
	ret.Format = "JSON"
	ret.Charset = "utf-8"
	ret.SignType = "MD5"
	ret.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	//ret.Version = "2.0"
	ret.bizContent = simplejson.New()
	ret.Path = consts.TPPath
	if len(ret.Config.TPDomain) == 0 {
		ret.Config.TPDomain = consts.TPDomain
	}
	return ret
}
// 将Request编码成POST请求的Body
func (req *TradeCreateRequest) Encode() (string, error) {
	//加签
	req.bizContent.Set("out_order_no", req.OutOrderNo)
	req.bizContent.Set("uid", req.Uid)
	req.bizContent.Set("uid_type", req.UidType)
	req.bizContent.Set("merchant_id", req.MerchantId)
	req.bizContent.Set("total_amount", req.TotalAmount)
	req.bizContent.Set("currency", req.Currency)
	req.bizContent.Set("subject", req.Subject)
	req.bizContent.Set("body", req.Body)
	req.bizContent.Set("product_code", req.ProductCode)
	req.bizContent.Set("payment_type", req.PaymentType)
	req.bizContent.Set("trade_time", req.TradeTime)
	req.bizContent.Set("valid_time", req.ValidTime)
	req.bizContent.Set("notify_url", req.NotifyUrl)
	req.bizContent.Set("service_fee", req.ServiceFee)
	req.bizContent.Set("risk_info", req.RiskInfo)

	// Json encode
	bizContentBytes, err := req.bizContent.Encode()
	if err != nil {
		util.Debug("TradeCreateRequest Encode bizContent.Encode err: %s, bizContent %s\n", err, *req.bizContent)
		return "", util.Wrap(err, "TradeCreateRequest Encode failed when [bizContent.Encode()]")
	}

	signParams := make(map[string]interface{})
	signParams["app_id"] = req.Config.AppId
	signParams["method"] = req.Method
	signParams["format"] = req.Format
	signParams["charset"] = req.Charset
	signParams["sign_type"] = req.SignType
	signParams["timestamp"] = req.Timestamp
	signParams["version"] = req.Version
	signParams["biz_content"] = string(bizContentBytes)

	sign := util.BuildMd5WithSalt(signParams, req.Config.AppSecret)
	// URL Encode
	values := url.Values{}
	values.Set("app_id", req.Config.AppId)
	values.Set("method", req.Method)
	values.Set("format", req.Format)
	values.Set("charset", req.Charset)
	values.Set("sign_type", req.SignType)
	values.Set("sign", sign)
	values.Set("timestamp", req.Timestamp)
	values.Set("version", req.Version)
	values.Set("biz_content", string(bizContentBytes))

	return values.Encode(), nil
}
// 生成此次请求logid
func (req *TradeCreateRequest) GetLogId() string {
	return fmt.Sprintf("%s_%s_%s_%s", req.AppId, req.MerchantId, req.OutOrderNo, req.Timestamp)
}

// 获取请求url地址
func (req *TradeCreateRequest) GetUrl() string {
	return req.Config.TPDomain + "/" + req.Path
}


//*********TradeCreateResponse*********//
// 预下单接口响应
type TradeCreateResponse struct {
	Data    *simplejson.Json
	TradeNo string `json:"trade_no"`
	URL     string `json:"url"`
	req     *TradeCreateRequest
}

// 初始化预下单响应
func NewTradeCreateResponse(req *TradeCreateRequest) *TradeCreateResponse {
	ret := new(TradeCreateResponse)
	ret.Data = simplejson.New()
	ret.req = req
	return ret
}
// 设置原始响应
func (resp *TradeCreateResponse) SetData(data *simplejson.Json) {
	resp.Data = data
}
// 将响应json数据反序列化为对应接口
func (resp *TradeCreateResponse) Decode() error {
	var respBytes []byte
	var err error
	// 走网关的接口拿到的参数在response里
	// 二维码接口拿到的参数在data里
	switch resp.Data.Get("response").Interface() {
	case nil:
		respBytes, err = resp.Data.Get("data").Encode()
	default:
		respBytes, err = resp.Data.Get("response").Encode()
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBytes, resp); err != nil {
		return err
	}
	return nil
}
