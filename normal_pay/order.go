package normal_pay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/v3/pay/partner/transactions/jsapi", bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := tools.Signature("v3/pay/partner/transactions/jsapi", string(dataJsonByte), string(keyByte), c.Config.SpMchID, c.Config.SerialNo)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "WECHATPAY2-SHA256-RSA2048 "+signature)
	client := http.Client{}
	resp, err := client.Do(req)
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
