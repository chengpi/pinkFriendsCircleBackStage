package trade

import (
	"context"
	"smallRoutine/loveta/models/tt_pay/util"
	"smallRoutine/loveta/models/tt_pay"
	"smallRoutine/loveta/models/tt_pay/consts"
	"fmt"
	"strconv"
	"time"
	"reflect"
)

// 预下单接口
func TradeCreate(ctx context.Context, req *TradeCreateRequest) (*TradeCreateResponse, error) {
	resp := NewTradeCreateResponse(req)
	// 1.0需要与财经后端通信取得"trade_no"
	if req.Version == "1.0" || req.AppletVersion == "2.0+" || req.AppletVersion == "1.0" {
		// 查验1.0参数
		if err := req.checkParams1_0(); err != nil {
			return nil, err
		}
		err := tt_pay.Execute(ctx, req.TPClientTimeoutMs, req, resp)
		// 当出现请求失败错误时，不封装
		fmt.Println(reflect.TypeOf(err))
		//switch err.(type) {
		//case error:
		//	fmt.Println("error")
		//case *util.Error:
		//	fmt.Println("*util.Error")
		//}

		_, ok := err.(*util.Error)
		//fmt.Println(ok)
		if ok {
			return nil, err
		}
		if err != nil {
			return nil, util.Wrap(err, "TradeCreate failed when [Execute]")
		}
	}
	if req.AppletVersion == "2.0" || req.AppletVersion == "2.0+" {
		// 查验2.0参数
		if err := req.checkParams2_0(); err != nil {
			return nil, err
		}
	}
	return resp, nil
}
// 1.0版小程序参数查验
func (req *TradeCreateRequest) checkParams1_0() error {
	if err := util.CheckAppId(req.AppId); err != nil {
		return err
	}

	if err := util.CheckFormat(req.Format); err != nil {
		return err
	}

	if err := util.CheckCharset(req.Charset); err != nil {
		return err
	}

	if err := util.CheckSignType(req.SignType); err != nil {
		return err
	}

	if err := util.CheckTimeStamp(req.Timestamp); err != nil {
		return err
	}

	if err := util.CheckVersion(req.Version); err != nil {
		return err
	}

	if err := util.CheckBizContent(req.bizContent); err != nil {
		return err
	}

	if err := util.CheckOutOrderNo(req.OutOrderNo); err != nil {
		return err
	}

	if err := util.CheckUid(req.Uid); err != nil {
		return err
	}

	if err := util.CheckMerchantId(req.MerchantId); err != nil {
		return err
	}

	if err := util.CheckTotalAmount(req.TotalAmount); err != nil {
		return err
	}

	if err := util.CheckCurrency(req.Currency); err != nil {
		return err
	}

	if err := util.CheckSubject(req.Subject); err != nil {
		return err
	}

	if err := util.CheckBody(req.Body); err != nil {
		return err
	}

	if err := util.CheckTradeTime(req.TradeTime); err != nil {
		return err
	}

	if err := util.CheckNotifyUrl(req.NotifyUrl); err != nil {
		return err
	}

	if err := util.CheckRiskInfo(req.RiskInfo); err != nil {
		return err
	}

	return nil
}

// 2.0版小程序参数查验
func (req *TradeCreateRequest) checkParams2_0() error {
	if err := util.CheckAppId(req.AppId); err != nil {
		return err
	}

	if err := util.CheckFormat(req.Format); err != nil {
		return err
	}

	if err := util.CheckCharset(req.Charset); err != nil {
		return err
	}

	if err := util.CheckSignType(req.SignType); err != nil {
		return err
	}

	if err := util.CheckTimeStamp(req.Timestamp); err != nil {
		return err
	}

	if err := util.CheckVersion(req.Version); err != nil {
		return err
	}

	if err := util.CheckBizContent(req.bizContent); err != nil {
		return err
	}

	if err := util.CheckOutOrderNo(req.OutOrderNo); err != nil {
		return err
	}

	if err := util.CheckUid(req.Uid); err != nil {
		return err
	}

	if err := util.CheckMerchantId(req.MerchantId); err != nil {
		return err
	}

	if err := util.CheckTotalAmount(req.TotalAmount); err != nil {
		return err
	}

	if err := util.CheckCurrency(req.Currency); err != nil {
		return err
	}

	if err := util.CheckSubject(req.Subject); err != nil {
		return err
	}

	if err := util.CheckBody(req.Body); err != nil {
		return err
	}

	if err := util.CheckTradeTime(req.TradeTime); err != nil {
		return err
	}

	if err := util.CheckNotifyUrl(req.NotifyUrl); err != nil {
		return err
	}

	if err := util.CheckRiskInfo(req.RiskInfo); err != nil {
		return err
	}

	if err := util.CheckProductCode(req.ProductCode); err != nil {
		return err
	}

	if err := util.CheckPaymentType(req.PaymentType); err != nil {
		return err
	}

	if err := util.CheckCashDeskTradeType(req.TradeType); err != nil {
		return err
	}

	if err := util.CheckValidTime(req.ValidTime); err != nil {
		return err
	}

	return nil
}


// 返回拉起小程序收银台的参数, json字符串
func (resp *TradeCreateResponse) GetCashdeskAppletParams() (string, error) {
	returnMap := make(map[string]string)
	switch resp.req.AppletVersion {
	case "1.0":
		returnJson, err := resp.getAppletParams1_0()
		if err != nil {
			return "", util.Wrap(err, "GetCashdeskAppletParams failed when[getAppletParams1_0()]")
		}
		returnMap["1.0"] = returnJson
	case "2.0":
		returnJson, err := resp.getAppletParams2_0()
		if err != nil {
			return "", util.Wrap(err, "GetCashdeskAppletParams failed when[getAppletParams2_0()]")
		}
		returnMap["2.0"] = returnJson
	case "2.0+":
		returnJson1_0, err := resp.getAppletParams1_0()
		if err != nil {
			return "", util.Wrap(err, "GetCashdeskAppletParams failed when[getAppletParams1_0()]")
		}
		returnJson2_0, err := resp.getAppletParams2_0()
		if err != nil {
			return "", util.Wrap(err, "GetCashdeskAppletParams failed when[getAppletParams2_0()]")
		}
		returnMap["1.0"] = returnJson1_0
		returnMap["2.0"] = returnJson2_0
	default:
		return "", fmt.Errorf(util.ErrorFormat, "AppletVerion", "AppletVersion can only be 1.0, 2.0 or 2.0+")
	}
	returnJson, err := util.JsonMarshal(returnMap)
	if err != nil {
		return "", util.Wrap(err, "GetCashdeskAppletParams failed when[JsonMarshal()]")
	}
	return returnJson, nil
}

// 返回拉起二维码收银台的参数，URL
func (resp *TradeCreateResponse) GetCashdeskQRParams() (string, error) {
	return resp.URL, nil
}

// 小程序1.0参数
func (resp *TradeCreateResponse) getAppletParams1_0() (string, error) {
	appletParams := make(map[string]interface{})

	appletParams["app_id"] = resp.req.AppId
	appletParams["sign_type"] = resp.req.SignType
	appletParams["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	appletParams["trade_no"] = resp.TradeNo
	appletParams["merchant_id"] = resp.req.MerchantId
	appletParams["uid"] = resp.req.Uid
	appletParams["total_amount"] = resp.req.TotalAmount
	appletParams["params"] = resp.req.Params

	appletParams["sign"] = util.BuildMd5WithSalt(appletParams, resp.req.AppSecret)

	appletParams["method"] = consts.MethodTradeConfirm // 方法要改为请求confirm
	appletParams["pay_type"] = resp.req.PaymentType1_0 // pay_type
	appletParams["pay_channel"] = resp.req.PayChannel
	appletParams["risk_info"] = resp.req.RiskInfo
	//if resp.req.ReturnUrl != "" {
	//	appletParams["return_url"] = resp.req.ReturnUrl
	//}
	//if resp.req.ShowURL != "" {
	//	appletParams["show_url"] = resp.req.ShowURL
	//}

	returnParams, err := util.JsonMarshal(appletParams)
	if err != nil {
		return "", util.Wrap(err, "getAppletParams1_0 failed when [JsonMarshal()]")
	}

	return returnParams, nil
}

// 小程序2.0参数
func (resp *TradeCreateResponse) getAppletParams2_0() (string, error) {
	cashDeskParams := make(map[string]interface{})

	cashDeskParams["app_id"] = resp.req.AppId
	cashDeskParams["sign_type"] = resp.req.SignType
	cashDeskParams["merchant_id"] = resp.req.MerchantId
	if resp.req.Uid != "" {
		cashDeskParams["uid"] = resp.req.Uid
	}
	if resp.req.OutOrderNo != "" {
		cashDeskParams["out_order_no"] = resp.req.OutOrderNo
	}
	cashDeskParams["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	cashDeskParams["total_amount"] = strconv.Itoa(resp.req.TotalAmount)
	if resp.req.NotifyUrl != "" {
		cashDeskParams["notify_url"] = resp.req.NotifyUrl
	}
	if resp.req.TradeType != "" {
		cashDeskParams["trade_type"] = resp.req.TradeType
	}
	if resp.req.ProductCode != "" {
		cashDeskParams["product_code"] = resp.req.ProductCode
	}
	if resp.req.PaymentType != "" {
		cashDeskParams["payment_type"] = resp.req.PaymentType
	}
	if resp.req.Subject != "" {
		cashDeskParams["subject"] = resp.req.Subject
	}
	if resp.req.Body != "" {
		cashDeskParams["body"] = resp.req.Body
	}
	if resp.req.TradeTime != "" {
		cashDeskParams["trade_time"] = resp.req.TradeTime
	}
	if resp.req.ValidTime != "" {
		cashDeskParams["valid_time"] = resp.req.ValidTime
	}
	if resp.req.Currency != "" {
		cashDeskParams["currency"] = resp.req.Currency
	}
	if resp.req.Version != "" {
		cashDeskParams["version"] = resp.req.Version
	}
	if resp.req.AlipayUrl != "" {
		cashDeskParams["alipay_url"] = resp.req.AlipayUrl
	}
	if resp.req.WxUrl != "" {
		cashDeskParams["wx_url"] = resp.req.WxUrl
	}
	if resp.req.WxType != "" {
		cashDeskParams["wx_type"] = resp.req.WxType
	}
	cashDeskParams["sign"] = util.BuildMd5WithSalt(cashDeskParams, resp.req.AppSecret)
	if resp.req.RiskInfo != "" {
		cashDeskParams["risk_info"] = resp.req.RiskInfo
	}

	returnParams, err := util.JsonMarshal(cashDeskParams)
	if err != nil {
		return "", util.Wrap(err, "getAppletParams2_0 failed when [JsonMarshal()]")
	}

	return returnParams, nil
}