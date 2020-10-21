package refund

import (
	"encoding/json"
	"io/ioutil"
	"log"

	weixin_shop_pay "github.com/pingyeaa/weixin-shop-pay"
	"github.com/pingyeaa/weixin-shop-pay/tools"
)

// Refund 退款
type Refund struct {
	Config *weixin_shop_pay.Config
}

// Apply 申请退款
func (c *Refund) Apply(p *weixin_shop_pay.RefundApply) (*ApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/refunds/apply"
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
	var output ApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyResp 下单返回数据
type ApplyResp struct {
	RefundID    string `json:"refund_id"`     // 微信退款单号
	OutRefundNo string `json:"out_refund_no"` // 商户退款单号
	CreateTime  string `json:"create_time"`   // 退款创建时间
	Amount      struct {
		Refund         int    `json:"refund"`          // 退款金额
		PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
		DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
		Currency       string `json:"currency"`        // 退款币种
	} `json:"amount"` // 订单金额
	PromotionDetail []struct {
		PromotionID  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
	} `json:"promotion_detail"` // 优惠退款详情
}

// Query 退款查询
func (c *Refund) Query(p *weixin_shop_pay.RefundQuery) (*QueryResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/refunds/id/" + p.RefundID
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
	var output QueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryResp 查询退款单
type QueryResp struct {
	RefundID            string `json:"refund_id"`             // 微信退款单号
	OutRefundNo         string `json:"out_refund_no"`         // 商户退款单号
	TransactionID       string `json:"transaction_id"`        // 微信订单号
	OutTradeNo          string `json:"out_trade_no"`          // 商户订单号
	Channel             string `json:"channel"`               // 退款渠道
	UserReceivedAccount string `json:"user_received_account"` // 退款入账账号
	SuccessTime         string `json:"success_time"`          // 退款成功时间
	CreateTime          string `json:"create_time"`           // 退款创建时间
	Status              string `json:"status"`                // 退款状态
	Amount              struct {
		Refund         int    `json:"refund"`          // 退款金额
		PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
		DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
		Currency       string `json:"currency"`        // 退款币种
	} `json:"amount"` // 订单金额
	PromotionDetail []struct {
		PromotionID  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
	} `json:"promotion_detail"` // 优惠退款详情
}
