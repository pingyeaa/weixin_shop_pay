package normal_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"weixin_shop_pay"
	"weixin_shop_pay/tools"
)

// Order 下单
func (c *NormalPay) Order(p *weixin_shop_pay.OrderParams) (*OrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/pay/partner/transactions/jsapi"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte, c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output OrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// OrderResp 下单返回数据
type OrderResp struct {
	PrepayID string `json:"prepay_id"`
}
