package ecommerce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"
	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Ecommerce 二级商户进件
type Ecommerce struct {
	Config *config.Config
}

// Apply 二级商户进件
func (c *Ecommerce) Apply(p *params.EcommerceApply) (*params.EcommerceApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/applyments"
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
	var output params.EcommerceApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyQuery 二级商户进件查询
func (c *Ecommerce) ApplyQuery(p *params.EcommerceApplyQuery) (*params.EcommerceApplyQueryResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/applyments"
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
	var output params.EcommerceApplyQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
