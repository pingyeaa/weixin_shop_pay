package balance

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"
	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Balance 余额
type Balance struct {
	Config *config.Config
}

// SubMch 二级商户余额查询
func (c *Balance) SubMch(p *params.BalanceSubMch) (*params.BalanceSubMchResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/fund/balance/" + p.SubMchid
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

	// 赋值返回
	log.Println(string(respData))
	var output params.BalanceSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchDate 二级商户余额查询
func (c *Balance) SubMchDate(p *params.BalanceSubMch) (*params.BalanceSubMchResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/fund/enddaybalance/" + p.SubMchid + "?date=" + p.Date
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

	// 赋值返回
	log.Println(string(respData))
	var output params.BalanceSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
