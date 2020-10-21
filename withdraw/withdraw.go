package withdraw

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"

	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Withdraw 普通支付
type Withdraw struct {
	Config *config.Config
}

// SubMch 二级商户余额提现
func (c *Withdraw) SubMch(p *params.WithdrawSubMch) (*params.WithdrawSubMchResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/withdraw"
	resp, err := tools.GetRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output params.WithdrawSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchQuery 二级商户余额查询
func (c *Withdraw) SubMchQuery(p *params.WithdrawSubMchQuery) (*params.WithdrawSubMchQueryResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/withdraw/" + p.WithdrawID
	resp, err := tools.GetRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output params.WithdrawSubMchQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
