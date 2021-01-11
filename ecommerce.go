package weixin_shop_pay

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

// Ecommerce 二级商户进件
type Ecommerce struct {
	Config *Config
}

// Apply 二级商户进件
func (c *Ecommerce) Apply(p *EcommerceApply) (*EcommerceApplyResp, error) {

	// 敏感信息加密
	var err error
	p.IDCardInfo.IDCardName, err = tool.Encrypt(p.IDCardInfo.IDCardName, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.IDCardInfo.IDCardNumber, err = tool.Encrypt(p.IDCardInfo.IDCardNumber, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactEmail, err = tool.Encrypt(p.ContactInfo.ContactEmail, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactIDCardNumber, err = tool.Encrypt(p.ContactInfo.ContactIDCardNumber, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.MobilePhone, err = tool.Encrypt(p.ContactInfo.MobilePhone, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactName, err = tool.Encrypt(p.ContactInfo.ContactName, c.Config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}

	if p.AccountInfo != nil {
		p.AccountInfo.AccountNumber, err = tool.Encrypt(p.AccountInfo.AccountNumber, c.Config.PlatformPublicKey)
		if err != nil {
			return nil, err
		}
		p.AccountInfo.AccountName, err = tool.Encrypt(p.AccountInfo.AccountName, c.Config.PlatformPublicKey)
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
	resp, err := tool.PostRequest(c.Config, urlPath, dataJsonByte)
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
		return nil, errors.New(string(respData))
	}

	log.Println(string(respData))
	var output EcommerceApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyQuery 二级商户进件查询
func (c *Ecommerce) ApplyQuery(p *EcommerceApplyQuery) (*EcommerceApplyQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/applyments/" + p.ApplymentID
	resp, err := tool.GetRequest(c.Config, urlPath)
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
		return nil, errors.New(string(respData))
	}

	log.Println(string(respData))
	var output EcommerceApplyQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ModifySettlement 修改结算账号
func (c *Ecommerce) ModifySettlement(p *EcommerceModifySettlement) error {
	// 加密银行卡号
	if p.AccountNumber != "" {
		AccountNumberMD, err := tool.Encrypt(p.AccountNumber, c.Config.PlatformPublicKey)
		if err != nil {
			return err
		}
		p.AccountNumber = AccountNumberMD
	}
	// 请求参数
	dataJsonByte, err := json.Marshal(EcommerceModifySettlementBody{
		AccountType:     p.AccountType,
		AccountBank:     p.AccountBank,
		BankAddressCode: p.BankAddressCode,
		AccountNumber:   p.AccountNumber,
	})
	if err != nil {
		return err
	}

	// 发起请求
	urlPath := "/v3/apply4sub/sub_merchants/" + p.SubMchid + "/modify-settlement"
	resp, err := tool.PostRequest(c.Config, urlPath, dataJsonByte)
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
		return errors.New(string(respData))
	}
	return nil
}

// QuerySettlement 查询结算信息
func (c *Ecommerce) QuerySettlement(p *EcommerceQuerySettlement) (*EcommerceQuerySettlementResp, error) {

	// 发起请求
	urlPath := "/v3/apply4sub/sub_merchants/" + p.SubMchid + "/settlement"
	resp, err := tool.GetRequest(c.Config, urlPath)
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
		return nil, errors.New(string(respData))
	}

	log.Println(string(respData))
	var output EcommerceQuerySettlementResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
