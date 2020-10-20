package profitsharing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"weixin_shop_pay"
	"weixin_shop_pay/tools"
)

// ProfitSharing 分账
type ProfitSharing struct {
	Config *weixin_shop_pay.Config
}

// ReceiversAdd 添加分账接收方
func (c *ProfitSharing) ReceiversAdd(p *weixin_shop_pay.QueryOrderParams) (*ReceiversAddResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/profitsharing/receivers/add" + p.TransactionID
	resp, err := tools.Request(urlPath, dataJsonByte, c.Config.KeyPath)
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
