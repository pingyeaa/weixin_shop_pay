package ecommerce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/pingyeaa/weixin_shop_pay/config"

	"github.com/pingyeaa/weixin_shop_pay/params"
	"github.com/pingyeaa/weixin_shop_pay/tools"
)

// Ecommerce 二级商户进件
type Ecommerce struct {
	Config *config.Config
}

// Apply 二级商户进件
func (c *Ecommerce) Apply(p *params.EcommerceApply) (*params.EcommerceApplyResp, error) {

	// 敏感信息加密
	var err error
	p.IDCardInfo.IDCardName, err = tools.Encrypt(p.IDCardInfo.IDCardName, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.IDCardInfo.IDCardNumber, err = tools.Encrypt(p.IDCardInfo.IDCardNumber, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactEmail, err = tools.Encrypt(p.ContactInfo.ContactEmail, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactIDCardNumber, err = tools.Encrypt(p.ContactInfo.ContactIDCardNumber, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.MobilePhone, err = tools.Encrypt(p.ContactInfo.MobilePhone, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactName, err = tools.Encrypt(p.ContactInfo.ContactName, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}

	if p.AccountInfo != nil {
		p.AccountInfo.AccountNumber, err = tools.Encrypt(p.AccountInfo.AccountNumber, c.Config.PlatformPublicKey)
		if err != nil {
			return nil, err
		}
		p.AccountInfo.AccountName, err = tools.Encrypt(p.AccountInfo.AccountName, c.Config.PlatformPublicKey)
		if err != nil {
			return nil, err
		}
	}

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/applyments/"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("二级进件接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.EcommerceApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyQuery 二级商户进件查询
func (c *Ecommerce) ApplyQuery(p *params.EcommerceApplyQuery) (*params.EcommerceApplyQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/applyments/" + p.ApplymentID
	resp, err := tools.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("二级进件查询接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.EcommerceApplyQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ModifySettlement 修改结算账号
func (c *Ecommerce) ModifySettlement(p *params.EcommerceModifySettlement) error {
	// 加密银行卡号
	if p.AccountNumber != "" {
		AccountNumberMD, err := tools.Encrypt(p.AccountNumber, c.Config.PlatformPublicKey)
		if err != nil {
			return err
		}
		p.AccountNumber = AccountNumberMD
	}
	// 请求参数
	dataJsonByte, err := json.Marshal(params.EcommerceModifySettlementBody{
		AccountType:     p.AccountType,
		AccountBank:     p.AccountBank,
		BankAddressCode: p.BankAddressCode,
		//BankName:        p.BankName,
		AccountNumber:   p.AccountNumber,
	})
	if err != nil {
		return err
	}

	// 发起请求
	urlPath := "/v3/apply4sub/sub_merchants/" + p.SubMchid + "/modify-settlement"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 验证接口是否错误
	if resp.StatusCode != 204 {
		return errors.New("修改结算账号接口请求异常：" + string(respData))
	}
	return nil
}

// QuerySettlement 查询结算信息
func (c *Ecommerce) QuerySettlement(p *params.EcommerceQuerySettlement) (*params.EcommerceQuerySettlementResp, error) {

	// 发起请求
	urlPath := "/v3/apply4sub/sub_merchants/" + p.SubMchid + "/settlement"
	resp, err := tools.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 验证接口是否错误
	if resp.StatusCode != 200 {
		return nil, errors.New("查询结算信息接口请求异常：" + string(respData))
	}

	log.Println(string(respData))
	var output params.EcommerceQuerySettlementResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
