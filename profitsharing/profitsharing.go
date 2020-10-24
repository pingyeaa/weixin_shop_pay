package profitsharing

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"

	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// ProfitSharing 分账
type ProfitSharing struct {
	Config *config.Config
}

// ReceiversAdd 添加分账接收方
func (c *ProfitSharing) ReceiversAdd(p *params.ProfitSharingReceiversAdd) (*params.ProfitSharingReceiversAddResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/receivers/add"
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
	var output params.ProfitSharingReceiversAddResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Apply 请求分账
func (c *ProfitSharing) Apply(p *params.ProfitSharingApply) (*params.ProfitSharingApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/orders"
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
	var output params.ProfitSharingApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Query 分账查询
func (c *ProfitSharing) Query(p *params.ProfitSharingQuery) (*params.ProfitSharingQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/orders"
	resp, err := tools.GetRequest(c.Config, urlPath)
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
	var output params.ProfitSharingQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// FinishOrder 完结分账
func (c *ProfitSharing) FinishOrder(p *params.ProfitSharingFinishOrder) (*params.ProfitSharingFinishOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/finish-order"
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
	var output params.ProfitSharingFinishOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
