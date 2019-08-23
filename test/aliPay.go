
/**
 * Created by GoLand
 * User: jackfu
 * Date: 2019-08-08
 * Time: 15:50
 */
package main

//import (
//"crypto"
//"crypto/md5"
//"crypto/rsa"
//"encoding/base64"
//"encoding/json"
//"errors"
//"fmt"
//"io"
//"sort"
//"strconv"
//"strings"
//"time"
//
//"code.qschou.com/openapi/refund_service/app/forbidden"
//
//"code.qschou.com/openapi/go-core/library/log/glog"
//
//"code.qschou.com/openapi/go-core/library/net/curl"
//"code.qschou.com/openapi/go-core/library/util"
//"code.qschou.com/openapi/refund_service/app/common"
//"code.qschou.com/openapi/refund_service/app/library/sign/grsa"
//"github.com/beevik/etree"
//"github.com/buger/jsonparser"
//"golang.org/x/net/html/charset"
//)
//
//type aliPay struct {
//	appId         string
//	appPid        string
//	appKey        string
//	signType      string
//	logCategory   string
//	appPrivateKey *rsa.PrivateKey // 应用私钥
//	aliPublicKey  *rsa.PublicKey  // 支付宝公钥
//}
//
//type AliPayRefund struct {
//	AlipayTradeRefundResponse struct {
//		Code                 string `json:"code"`
//		Msg                  string `json:"msg"`
//		SubMsg               string `json:"sub_msg"`
//		SubCode              string `json:"sub_code"`
//		TradeNo              string `json:"trade_no"`
//		OutTradeNo           string `json:"out_trade_no"`
//		BuyerLogonID         string `json:"buyer_logon_id"`
//		FundChange           string `json:"fund_change"`
//		RefundFee            string `json:"refund_fee"`
//		RefundCurrency       string `json:"refund_currency"`
//		GmtRefundPay         string `json:"gmt_refund_pay"`
//		RefundDetailItemList []struct {
//			FundChannel string  `json:"fund_channel"`
//			BankCode    string  `json:"bank_code"`
//			Amount      int     `json:"amount"`
//			RealAmount  float64 `json:"real_amount"`
//			FundType    string  `json:"fund_type"`
//		} `json:"refund_detail_item_list"`
//		StoreName               string `json:"store_name"`
//		BuyerUserID             string `json:"buyer_user_id"`
//		RefundPresetPaytoolList struct {
//			Amount         []float64 `json:"amount"`
//			AssertTypeCode string    `json:"assert_type_code"`
//		} `json:"refund_preset_paytool_list"`
//		RefundSettlementID           string `json:"refund_settlement_id"`
//		PresentRefundBuyerAmount     string `json:"present_refund_buyer_amount"`
//		PresentRefundDiscountAmount  string `json:"present_refund_discount_amount"`
//		PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount"`
//	} `json:"alipay_trade_refund_response"`
//	Sign string `json:"sign"`
//}
//
//func GetAliPayClient(business string) (*aliPay, error) {
//	var (
//		client = new(aliPay)
//		config = make(map[string]string)
//	)
//	if config = forbidden.GetPayConfig(business); config == nil {
//		return nil, errors.New("get " + business + " config failed")
//	}
//	client.appId = config["appId"]
//	client.appKey = config["appKey"]
//	client.appPid = config["pid"]
//	client.logCategory = "aliPayProxy_%s"
//	client.signType = common.AliPaySignType
//	if privateKey, err := grsa.ParsePKCS1PrivateKey(grsa.FormatPrivateKey(config["privateKey"])); err == nil {
//		client.appPrivateKey = privateKey
//	} else {
//		return nil, err
//	}
//	if publicKey, err := grsa.ParsePKCS1PublicKey(grsa.FormatPublicKey(config["publicKey"])); err == nil {
//		client.aliPublicKey = publicKey
//	} else {
//		return nil, err
//	}
//	return client, nil
//}
//
//// 请求支付宝接口
//func (a *aliPay) request(method string, bizParams map[string]string) (response []byte, err error) {
//	var (
//		sign       string
//		hash       crypto.Hash
//		params     = make(map[string]string)
//		bizContent []byte
//	)
//	if bizContent, err = json.Marshal(bizParams); err != nil {
//		return nil, err
//	}
//	params = map[string]string{
//		"app_id":      a.appId,
//		"method":      method,
//		"format":      "JSON",
//		"charset":     "utf-8",
//		"sign_type":   a.signType,
//		"timestamp":   time.Now().Format(util.TimeFormatDateTime),
//		"version":     "1.0",
//		"biz_content": string(bizContent),
//	}
//	if a.signType == common.SignTypeRSA2 {
//		hash = crypto.SHA256
//	} else {
//		hash = crypto.SHA1
//	}
//	if sign, err = a.genRsaSign(params, hash); err != nil {
//		return nil, err
//	}
//	params["sign"] = sign
//	response, err = curl.CommonReq.PostForm(common.AliPayDomain, params)
//	return
//}
//
//// 生成rsa参数签名
//func (a *aliPay) genRsaSign(params map[string]string, hash crypto.Hash) (sign string, err error) {
//	var (
//		sig        []byte
//		paramsStr  string
//		paramsList = make([]string, 0)
//	)
//	for k, v := range params {
//		paramsList = append(paramsList, k+"="+v)
//	}
//	sort.Strings(paramsList)
//	paramsStr = strings.Join(paramsList, "&")
//	if sig, err = grsa.SignPKCS1v15WithKey([]byte(paramsStr), a.appPrivateKey, hash); err != nil {
//		return "", err
//	}
//	sign = base64.StdEncoding.EncodeToString(sig)
//	return
//}
//
//// 验证数据返回
//func (a *aliPay) checkResponse(method string, response []byte) (err error) {
//	if a.aliPublicKey == nil {
//		return nil
//	}
//	var (
//		sign         string
//		content      []byte
//		signBytes    []byte
//		rootNodeName string
//	)
//	rootNodeName = strings.Replace(method, ".", "_", -1) + common.ResponseSuffix
//	if content, _, _, err = jsonparser.Get(response, rootNodeName); err != nil {
//		return err
//	}
//	if sign, err = jsonparser.GetString(response, "sign"); err != nil {
//		return err
//	}
//
//	if signBytes, err = base64.StdEncoding.DecodeString(sign); err != nil {
//		return err
//	}
//
//	if a.signType == common.SignTypeRSA2 {
//		err = grsa.VerifyPKCS1v15WithKey(content, signBytes, a.aliPublicKey, crypto.SHA256)
//	} else {
//		err = grsa.VerifyPKCS1v15WithKey(content, signBytes, a.aliPublicKey, crypto.SHA1)
//	}
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// 订单查询
//func (a *aliPay) TradeQuery(logIns glog.BaseLog, tradeNo, outTradeNo string) {
//	logCategory := fmt.Sprintf(a.logCategory, "TradeQuery")
//	var (
//		err      error
//		params   = make(map[string]string)
//		response []byte
//	)
//	params = map[string]string{
//		"trade_no":     tradeNo,
//		"out_trade_no": outTradeNo,
//	}
//	if response, err = a.request(common.AliPayTradeQuery, params); err != nil {
//		logIns.WarnMsg(logCategory, params).Msg(err.Error())
//		return
//	}
//	fmt.Println(string(response))
//	if err = a.checkResponse(common.AliPayTradeQuery, response); err != nil {
//		fmt.Println(err)
//		// logIns.WarnMsg(logCategory, map[string]interface{}{"params": params, "response": string(response)}).Msg(err.Error())
//		return
//	}
//	fmt.Println("success")
//}
//
//// 退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund/
//func (a *aliPay) Refund(logIns glog.BaseLog, params map[string]string) (responseData string, err error) {
//	logCategory := fmt.Sprintf(a.logCategory, "aliPayRefund")
//	//请求参数
//	requestParams := make(map[string]string)
//
//	fmt.Println(requestParams)
//	if "" == params["outTradeNo"] || "" ==params["tradeNo"] || "" == params["refundAmount"] || "" ==params["outRequestNo"] {
//		errMsg := "缺少参数"
//		fmt.Println(errMsg)
//		err = errors.New(errMsg)
//		return
//	}
//	//请求参数
//	requestParams = map[string]string{
//		"out_trade_no":  params["outTradeNo"],
//		"trade_no":      params["tradeNo"],
//		"refund_amount": params["refundAmount"],
//		"out_request_no": params["outRequestNo"],
//	}
//	response, err := a.request(common.AliPayRefund, requestParams)
//	if err != nil {
//		logIns.WarnMsg(logCategory, params).Msg(err.Error())
//		return
//	}
//	//接收响应参数至结构体
//	refundResponse := new(AliPayRefund)
//	err = json.Unmarshal(response, refundResponse)
//	if err != nil {
//		//logIns.WarnMsg(logCategory, params).Msg(err.Error())
//		return
//	}
//	//TODO:
//	if "TRADE_NOT_ALLOW_REFUND" == refundResponse.AlipayTradeRefundResponse.SubCode {
//
//	}
//
//	responseData = string(response)
//	return
//}
//
//// 退款查询接口 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query/
//func (a *aliPay) RefundQuery(logIns glog.BaseLog, tradeNo, outTradeNo, outRequestNo string) {
//	// TODO 完善返回值
//	logCategory := fmt.Sprintf(a.logCategory, "RefundQuery")
//	var (
//		err      error
//		params   = make(map[string]string)
//		response []byte
//	)
//	params = map[string]string{
//		"trade_no":       tradeNo,
//		"out_trade_no":   outTradeNo,
//		"out_request_no": outRequestNo,
//	}
//	if response, err = a.request(common.AliPayRefundQuery, params); err != nil {
//		logIns.WarnMsg(logCategory, params).Msg(err.Error())
//		return
//	}
//	if err = a.checkResponse(common.AliPayRefundQuery, response); err != nil {
//		logIns.WarnMsg(logCategory, map[string]interface{}{"params": params, "response": string(response)}).Msg(err.Error())
//		return
//	}
//	// TODO
//}
//
//// 老版订单查询接口
//func (a *aliPay) TradeQueryOld(logIns glog.BaseLog, tradeNo string) (result map[string]interface{}, err error) {
//	logCategory := fmt.Sprintf(a.logCategory, "TradeQueryOld")
//	// TODO 完善
//	var (
//		params     = make(map[string]string)
//		response   []byte
//		paramsStr  string
//		paramsList = make([]string, 0)
//	)
//	result = map[string]interface{}{
//		"is_paid":       false,
//		"refund_status": "",
//		"refund_all":    false,
//		"refund_amount": 0.00,
//		"total_amount":  0.00,
//	}
//	params = map[string]string{
//		"_input_charset": "utf-8",
//		"partner":        a.appPid,
//		"service":        "single_trade_query",
//		"out_trade_no":   tradeNo,
//	}
//	for k, v := range params {
//		paramsList = append(paramsList, k+"="+v)
//	}
//	sort.Strings(paramsList)
//	paramsStr = strings.Join(paramsList, "&") + a.appKey
//	fmt.Println(paramsStr)
//	params["sign"] = fmt.Sprintf("%x", md5.Sum([]byte(paramsStr)))
//	params["sign_type"] = "MD5"
//	if response, err = curl.CommonReq.PostForm(common.AliPayOldDomain, params); err != nil {
//		logIns.WarnMsg(logCategory, map[string]interface{}{"params": params, "response": string(response)}).Msg(err.Error())
//		return
//	}
//	doc := etree.NewDocument()
//	fmt.Println(string(response))
//	doc.ReadSettings.CharsetReader = func(char string, input io.Reader) (reader io.Reader, e error) {
//		return charset.NewReader(input, char)
//		//return input, nil
//	}
//	// TODO fix 字符集  recovery   字段不存在报panic
//	_ = doc.ReadFromBytes(response)
//	/*if err = doc.ReadFromString(string(response)); err != nil {
//		fmt.Println(err.Error())
//		logIns.WarnMsg(logCategory, map[string]interface{}{"params": params, "response": string(response)}).Msg(err.Error())
//		//return
//	}*/
//	if "T" != doc.SelectElement("alipay").SelectElement("is_success").Text() {
//		logIns.WarnMsg(logCategory, map[string]interface{}{"params": params, "response": string(response)}).Msg("query error")
//		return
//	}
//	trade := doc.SelectElement("alipay").SelectElement("response").SelectElement("trade")
//	fmt.Println(trade.SelectElement("gmt_payment").Text())
//	fmt.Println(trade.SelectElement("to_buyer_fee").Text())
//	if trade.SelectElement("gmt_payment").Text() != "" && trade.SelectElement("to_seller_fee").Text() != "" {
//		result["is_paid"] = true
//	}
//	result["refund_status"] = trade.SelectElement("refund_status").Text()
//	if trade.SelectElement("refund_fee").Text() == trade.SelectElement("price").Text() {
//		result["refund_all"] = true
//	}
//	if trade.SelectElement("refund_fee").Text() != "" {
//		result["refund_amount"], _ = strconv.ParseFloat(trade.SelectElement("refund_fee").Text(), 64)
//	}
//	if trade.SelectElement("total_fee").Text() != "" {
//		result["total_amount"], _ = strconv.ParseFloat(trade.SelectElement("total_fee").Text(), 64)
//	}
//	fmt.Println(result)
//	return result, nil
//}
