package refund

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"

	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Refund 退款
type Refund struct {
	Config *config.Config
}

// Apply 申请退款
func (c *Refund) Apply(p *params.RefundApply) (*params.RefundApplyResp, error) {

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

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("余额查询接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.RefundApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Query 退款查询
func (c *Refund) Query(p *params.RefundQuery) (*params.RefundQueryResp, error) {

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

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("余额查询接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.RefundQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
