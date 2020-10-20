package weixin_shop_pay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// OrderParams 下单参数
type OrderParams struct {
	SpAppID     string `json:"sp_appid"`     // 服务商公众号ID
	SpMchID     string `json:"sp_mchid"`     // 服务商户号
	SubAppID    string `json:"sub_appid"`    // 二级商户公众号ID
	SubMchID    string `json:"sub_mchid"`    // 二级商户号
	Description string `json:"description"`  // 商品描述
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号
	TimeExpire  string `json:"time_expire"`  // 交易结束时间
	Attach      string `json:"attach"`       // 附加数据
	NotifyURL   string `json:"notify_url"`   // 通知地址
	GoodsTag    string `json:"goods_tag"`    // 订单优惠标记
	SettleInfo  struct {
		ProfitSharing bool  `json:"profit_sharing"` // 是否制定分账
		SubsidyAmount int64 `json:"subsidy_amount"` // 补差金额
	} `json:"settle_info"` // 结算信息
	Amount    *Amount    `json:"amount"`     // 订单金额
	Payer     *Payer     `json:"payer"`      // 支付者
	Detail    *Detail    `json:"detail"`     // 优惠功能
	SceneInfo *SceneInfo `json:"scene_info"` // 场景信息
}

// Order 下单
func (c *Client) Order(p *OrderParams) (*OrderResp, error) {

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
	signature, err := Signature("v3/pay/partner/transactions/jsapi", string(dataJsonByte), string(keyByte), c.Config.SpMchID, c.Config.SerialNo)
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
