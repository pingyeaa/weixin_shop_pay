package weixin_shop_pay

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

// Withdraw 普通支付
type Withdraw struct {
	Config *Config
}

// SubMch 二级商户余额提现
func (c *Withdraw) SubMch(p *WithdrawSubMch) (*WithdrawSubMchResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	// 发起请求
	urlPath := "/v3/ecommerce/fund/withdraw"
	resp, err := tool.PostRequest(c.Config, urlPath, dataJsonByte)
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
		return nil, errors.New(string(respData))
	}

	log.Println(string(respData))
	var output WithdrawSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchQuery 二级商户提现状态查询
func (c *Withdraw) SubMchQuery(p *WithdrawSubMchQuery) (*WithdrawSubMchQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/fund/withdraw/" + p.WithdrawID + "?sub_mchid=" + p.SubMchid
	resp, err := tool.GetRequest(c.Config, urlPath)
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
		return nil, errors.New(string(respData))
	}

	log.Println(string(respData))
	var output WithdrawSubMchQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
