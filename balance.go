package weixin_shop_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Balance 余额
type Balance struct {
	client *Client
}

// SubMch 二级商户余额查询
func (t *Balance) SubMch(p *BalanceSubMch) (*BalanceSubMchResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/fund/balance/" + p.SubMchid
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	// 赋值返回
	log.Println(string(respData))
	var output BalanceSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchData二级商户按日期余额查询 https://api.mch.weixin.qq.com/v3/ecommerce/fund/enddaybalance/1222212222?date=2019-08-17
func (c *Balance) SubMchData(p *params.BalanceSubMch) (*params.BalanceSubMchResp, error) {

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
		return nil, errors.New("余额日期查询接口请求异常：" + string(respData))
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

