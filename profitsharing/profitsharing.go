package profitsharing

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin-shop-pay/config"

	"github.com/pingyeaa/weixin-shop-pay/params"

	"github.com/pingyeaa/weixin-shop-pay/tools"
)

// ProfitSharing 分账
type ProfitSharing struct {
	Config *config.Config
}

// ReceiversAdd 添加分账接收方
func (c *ProfitSharing) ReceiversAdd(p *params.PayQueryOrder) (*ReceiversAddResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/profitsharing/receivers/add" + p.TransactionID
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
	var output ReceiversAddResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryOrderResp 订单查询返回参数
type ReceiversAddResp struct {
	Type    string `json:"type"`    // 接收方类型
	Account string `json:"account"` // 接收方账号
}

// Apply 请求分账
func (c *ProfitSharing) Apply(p *params.ProfitSharingApply) (*ApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/profitsharing/orders"
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

// ApplyResp .
type ApplyResp struct {
	Type    string `json:"type"`    // 接收方类型
	Account string `json:"account"` // 接收方账号
}

// Query 分账查询
func (c *ProfitSharing) Query(p *params.ProfitSharingQuery) (*QueryResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/profitsharing/orders"
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

// QueryResp .
type QueryResp struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderID       string `json:"order_id"`       // 微信分账单号
	Status        string `json:"status"`         // 分账状态
	Receivers     []struct {
		ReceiverMchid   string `json:"receiver_mchid"`   // 分账接收商户号
		Amount          int    `json:"amount"`           // 分账金额
		Description     string `json:"description"`      // 分账描述
		Result          string `json:"result"`           // 分账结果
		FinishTime      string `json:"finish_time"`      // 完成时间
		FailReason      string `json:"fail_reason"`      // 分账失败原因
		Type            string `json:"type"`             // 分账接收方类型
		ReceiverAccount string `json:"receiver_account"` // 分账接收方账号
	} `json:"receivers"` // 分账接收方列表
	CloseReason       string `json:"close_reason"`       // 关单原因
	FinishAmount      int    `json:"finish_amount"`      // 分账完结金额
	FinishDescription string `json:"finish_description"` // 分账完结描述
}

// FinishOrder 完结分账
func (c *ProfitSharing) FinishOrder(p *params.ProfitSharingFinishOrder) (*FinishOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/profitsharing/finish-order"
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
	var output FinishOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// FinishOrderResp .
type FinishOrderResp struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderID       string `json:"order_id"`       // 微信分账单号
}
