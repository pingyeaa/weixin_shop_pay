package weixin_shop_pay

import (
	"log"
	"testing"
)

func TestCommon_ImageUpload(t *testing.T) {
	common := Common{Config: &Config{
		SpAppID:  "",
		SpMchID:  "",
		KeyPath:  "",
		SerialNo: "",
	}}
	resp, err := common.ImageUpload(&CommonImageUpload{
		FilePath: "/Users/zhangyinuo/Golang/weixin-shop-pay/test.jpg",
	})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
