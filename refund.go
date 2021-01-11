package weixin_shop_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Refund 退款
type Refund struct {
	Config *Config
}

// Apply 申请退款
func (c *Refund) Apply(p *RefundApply) (*RefundApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/refunds/apply"
	resp, err := tool.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respData))
	var output RefundApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Query 退款查询
func (c *Refund) Query(p *RefundQuery) (*RefundQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/refunds/id/" + p.RefundID
	resp, err := tool.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respData))
	var output RefundQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// QueryByRefundNo 退款查询
func (c *Refund) QueryByRefundNo(p *RefundQueryByRefundNo) (*RefundQueryResp, error) {

	// 发起请求
	// /v3/ecommerce/refunds/out-refund-no/{out_refund_no}
	urlPath := "/v3/ecommerce/refunds/out-refund-no/" + p.OutRefundNo + "?sub_mchid=" + p.SubMchid

	resp, err := tool.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respData))
	var output RefundQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
