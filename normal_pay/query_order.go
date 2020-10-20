package normal_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"weixin_shop_pay"
	"weixin_shop_pay/tools"
)

// QueryOrder 查询订单
func (c *NormalPay) QueryOrder(p *weixin_shop_pay.QueryOrderParams) (*QueryOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/pay/partner/transactions/id/" + p.TransactionID
	resp, err := tools.Request(urlPath, dataJsonByte, c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output QueryOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderResp 订单查询返回参数
type QueryOrderResp struct {
	SpAppID        string                     `json:"sp_appid"`         // 服务商公众号ID
	SpMchID        string                     `json:"sp_mchid"`         // 服务商户号
	SubAppID       string                     `json:"sub_appid"`        // 二级商户公众号ID
	SubMchID       string                     `json:"sub_mchid"`        // 二级商户号
	OutTradeNo     string                     `json:"out_trade_no"`     // 商户订单号
	TransactionID  string                     `json:"transaction_id"`   // 微信支付订单号
	TradeType      string                     `json:"trade_type"`       // 交易类型
	TradeState     string                     `json:"trade_state"`      // 交易状态
	TradeStateDesc string                     `json:"trade_state_desc"` // 交易状态描述
	BankType       string                     `json:"bank_type"`        // 付款银行
	Attach         string                     `json:"attach"`           // 附加数据
	SuccessTime    string                     `json:"success_time"`     // 付款完成时间
	Payer          *weixin_shop_pay.Payer     `json:"payer"`            // 支付者
	Amount         *weixin_shop_pay.Amount    `json:"amount"`           // 订单金额
	Detail         *weixin_shop_pay.Detail    `json:"detail"`           // 优惠功能
	SceneInfo      *weixin_shop_pay.SceneInfo `json:"scene_info"`       // 场景信息
}
