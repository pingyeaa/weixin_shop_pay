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
