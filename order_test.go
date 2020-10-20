package weixin_shop_pay

import (
	"log"
	"testing"
)

func TestClient_Order(t *testing.T) {
	client := &Client{}
	resp, err := client.Order(&OrderParams{
		SpAppID:     "",
		SpMchID:     "",
		SubAppID:    "",
		SubMchID:    "",
		Description: "",
		OutTradeNo:  "",
		TimeExpire:  "",
		Attach:      "",
		NotifyURL:   "",
		GoodsTag:    "",
		SettleInfo: struct {
			ProfitSharing bool  `json:"profit_sharing"`
			SubsidyAmount int64 `json:"subsidy_amount"`
		}{},
		Amount: struct {
			Total    int    `json:"total"`
			Currency string `json:"currency"`
		}{},
		Payer: struct {
			SpOpenid  string `json:"sp_openid"`
			SubOpenid string `json:"sub_openid"`
		}{},
	})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.PrepayID)
}
