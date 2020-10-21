package pay

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin-shop-pay/config"

	"github.com/pingyeaa/weixin-shop-pay/params"
	"github.com/pingyeaa/weixin-shop-pay/tools"
)

// Pay 普通支付
type Pay struct {
	Config *config.Config
}

// Order 下单
func (c *Pay) Order(p *params.PayOrder) (*OrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/pay/partner/transactions/jsapi"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output OrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// OrderResp 下单返回数据
type OrderResp struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	TradeType  string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty" json:"code_url,omitempty"`
	MwebUrl    string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty"`
	PrepayID   string `json:"prepay_id"`
}

// QueryOrder 查询订单
func (c *Pay) QueryOrder(p *params.PayQueryOrder) (*QueryOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/pay/partner/transactions/id/" + p.TransactionID
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
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
	SpAppID        string                    `json:"sp_appid"`         // 服务商公众号ID
	SpMchID        string                    `json:"sp_mchid"`         // 服务商户号
	SubAppID       string                    `json:"sub_appid"`        // 二级商户公众号ID
	SubMchID       string                    `json:"sub_mchid"`        // 二级商户号
	OutTradeNo     string                    `json:"out_trade_no"`     // 商户订单号
	TransactionID  string                    `json:"transaction_id"`   // 微信支付订单号
	TradeType      string                    `json:"trade_type"`       // 交易类型
	TradeState     string                    `json:"trade_state"`      // 交易状态
	TradeStateDesc string                    `json:"trade_state_desc"` // 交易状态描述
	BankType       string                    `json:"bank_type"`        // 付款银行
	Attach         string                    `json:"attach"`           // 附加数据
	SuccessTime    string                    `json:"success_time"`     // 付款完成时间
	Payer          *params.PayOrderPayer     `json:"payer"`            // 支付者
	Amount         *params.PayOrderAmount    `json:"amount"`           // 订单金额
	Detail         *params.PayOrderDetail    `json:"detail"`           // 优惠功能
	SceneInfo      *params.PayOrderSceneInfo `json:"scene_info"`       // 场景信息
}
