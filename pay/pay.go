package pay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"
	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Pay 普通支付
type Pay struct {
	Config *config.Config
}

// Order 下单
func (c *Pay) Order(p *params.PayOrder) (*params.PayOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/jsapi"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	log.Println("响应结果", string(respData))
	log.Println("响应头信息", resp.StatusCode, resp.Status)
	if resp.StatusCode != 200 {
		return nil, errors.New("下单接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.PayOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderTransaction 微信订单查询
func (c *Pay) QueryOrderTransaction(p *params.PayQueryOrderTransaction) (*params.PayQueryOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/id/" + p.TransactionID
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("余额查询接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.PayQueryOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderOutTradeNo 商户订单查询
func (c *Pay) QueryOrderOutTradeNo(p *params.PayQueryOrderOutTradeNo) (*params.PayQueryOrderResp, error) {

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/out-trade-no/" + p.OutTradeNo + fmt.Sprintf("?sp_mchid=%s&sub_mchid=%s", p.SpMchID, p.SubMchID)
	resp, err := tools.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("余额查询接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.PayQueryOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
