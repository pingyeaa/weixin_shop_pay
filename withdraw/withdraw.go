package withdraw

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"weixin_shop_pay"
	"weixin_shop_pay/tools"
)

// Withdraw 普通支付
type Withdraw struct {
	Config *weixin_shop_pay.Config
}

// SubMch 二级商户余额提现
func (c *Withdraw) SubMch(p *weixin_shop_pay.WithdrawSubMch) (*SubMchResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/withdraw"
	resp, err := tools.GetRequest(c.Config, urlPath, dataJsonByte, c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output SubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchResp .
type SubMchResp struct {
	SubMchid     string `json:"sub_mchid"`      // 二级商户号
	WithdrawID   string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo string `json:"out_request_no"` // 商户提现单号s
}

// SubMchQuery 二级商户余额查询
func (c *Withdraw) SubMchQuery(p *weixin_shop_pay.WithdrawSubMchQuery) (*SubMchQueryResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/withdraw/" + p.WithdrawID
	resp, err := tools.GetRequest(c.Config, urlPath, dataJsonByte, c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output SubMchQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchQueryResp .
type SubMchQueryResp struct {
	SubMchid     string `json:"sub_mchid"`      // 二级商户号
	SpMchid      string `json:"sp_mchid"`       // 电商平台商户号
	Status       string `json:"status"`         // 提现单状态
	WithdrawID   string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo string `json:"out_request_no"` // 商户提现单号
	Amount       int    `json:"amount"`         // 提现金额
	CreateTime   string `json:"create_time"`    // 发起提现时间
	UpdateTime   string `json:"update_time"`    // 提现状态更新时间
	Reason       string `json:"reason"`         // 失败原因
	Remark       string `json:"remark"`         // 提现备注
	BankMemo     string `json:"bank_memo"`      // 银行附言
}
