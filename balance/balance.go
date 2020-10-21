package balance

import (
	"encoding/json"
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

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/balance/" + p.SubMchid
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
	var output params.BalanceSubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
