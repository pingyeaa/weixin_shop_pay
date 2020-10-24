package common

import (
	"log"
	"testing"

	"github.com/pingyeaa/weixin_shop_pay/config"
	"github.com/pingyeaa/weixin_shop_pay/params"
)

func TestCommon_ImageUpload(t *testing.T) {
	common := Common{Config: &config.Config{
		SpAppID:  "",
		SpMchID:  "",
		KeyPath:  "",
		SerialNo: "",
	}}
	resp, err := common.ImageUpload(&params.CommonImageUpload{
		FilePath: "/Users/zhangyinuo/Golang/weixin-shop-pay/test.jpg",
	})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
