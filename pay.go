package weixin_shop_pay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Pay 普通支付
type Pay struct {
	client *Client
}

// Order 下单
func (t *Pay) Order(p *PayOrder) (*PayOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/jsapi"
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
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
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output PayOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderTransaction 微信订单查询
func (t *Pay) QueryOrderTransaction(p *PayQueryOrderTransaction) (*PayQueryOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/id/" + p.TransactionID
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output PayQueryOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderOutTradeNo 商户订单查询
func (t *Pay) QueryOrderOutTradeNo(p *PayQueryOrderOutTradeNo) (*PayQueryOrderResp, error) {

	// 发起请求
	urlPath := "/v3/pay/partner/transactions/out-trade-no/" + p.OutTradeNo + fmt.Sprintf("?sp_mchid=%s&sub_mchid=%s", p.SpMchID, p.SubMchID)
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output PayQueryOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
