package ecommerce

import (
	"log"
	"testing"

	"github.com/pingyeaa/weixin_shop_pay/config"
	"github.com/pingyeaa/weixin_shop_pay/params"
)

func TestEcommerce_Apply(t *testing.T) {
	ecommerce := Ecommerce{Config: &config.Config{
		SpAppID: "",
	}}
	resp, err := ecommerce.Apply(&params.EcommerceApply{
		OutRequestNo:         "",
		OrganizationType:     "",
		BusinessLicenseInfo:  nil,
		OrganizationCertInfo: nil,
		IDDocType:            "",
		IDCardInfo:           nil,
		IDDocInfo:            nil,
		NeedAccountInfo:      false,
		AccountInfo:          nil,
		ContactInfo:          nil,
		SalesSceneInfo:       nil,
		MerchantShortname:    "",
		Qualifications:       "",
		BusinessAdditionPics: "",
		BusinessAdditionDesc: "",
	})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}

func TestEcommerce_ApplyQuery(t *testing.T) {
	ecommerce := Ecommerce{Config: &config.Config{
		SpAppID:  "",
		SpMchID:  "1603361319",
		KeyPath:  "/Users/zhangyinuo/Golang/hdzs-api-go/cert/service_provider/apiclient_key.pem",
		SerialNo: "6F7BFF8A4B560251EA5FA28FE961A8ABA4DDA063",
	}}
	resp, err := ecommerce.ApplyQuery(&params.EcommerceApplyQuery{ApplymentID: "123"})
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
