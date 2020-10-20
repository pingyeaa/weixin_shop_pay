package weixin_shop_pay

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient(&Config{
		SpAppID:  "",
		SpMchID:  "",
		KeyPath:  "",
		SerialNo: "",
	})
	_, err := client.NormalPay().Order(&OrderParams{
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
		Amount:    nil,
		Payer:     nil,
		Detail:    nil,
		SceneInfo: nil,
	})
	if err != nil {
		t.Error(err)
	}

}
