package weixin_shop_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Withdraw 普通支付
type Withdraw struct {
	client *Client
}

// SubMch 二级商户余额提现
func (t *Withdraw) SubMch(p *WithdrawSubMch) (*WithdrawSubMchResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	// 发起请求
	urlPath := "/v3/ecommerce/fund/withdraw"
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
	var output WithdrawSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchQuery 二级商户提现状态查询
func (t *Withdraw) SubMchQuery(subMchID string, withdrawID string) (*WithdrawSubMchQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/fund/withdraw/" + withdrawID + "?sub_mchid=" + subMchID
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
	var output WithdrawSubMchQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
