package common

import (
	"log"
	"testing"

	"github.com/pingyeaa/weixin_shop_pay/params"

	"github.com/pingyeaa/weixin_shop_pay/config"
)

func TestCommon_ImageUpload(t *testing.T) {
	common := Common{Config: &config.Config{
		SpAppID:  "",
		SpMchID:  "1603361319",
		KeyPath:  "/Users/zhangyinuo/Golang/hdzs-api-go/cert/service_provider/apiclient_key.pem",
		SerialNo: "6F7BFF8A4B560251EA5FA28FE961A8ABA4DDA063",
	}}
	resp, err := common.ImageUpload(&params.CommonImageUpload{
		FilePath: "/Users/zhangyinuo/Golang/weixin-shop-pay/test.jpg",
	})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
