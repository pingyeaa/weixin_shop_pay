package weixin_shop_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Ecommerce 二级商户进件
type Ecommerce struct {
	client *Client
}

// Apply 二级商户进件
func (t *Ecommerce) Apply(p *EcommerceApply) (*EcommerceApplyResp, error) {

	// 敏感信息加密
	var err error
	p.IDCardInfo.IDCardName, err = tool.Encrypt(p.IDCardInfo.IDCardName, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.IDCardInfo.IDCardNumber, err = tool.Encrypt(p.IDCardInfo.IDCardNumber, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactEmail, err = tool.Encrypt(p.ContactInfo.ContactEmail, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactIDCardNumber, err = tool.Encrypt(p.ContactInfo.ContactIDCardNumber, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.MobilePhone, err = tool.Encrypt(p.ContactInfo.MobilePhone, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}
	p.ContactInfo.ContactName, err = tool.Encrypt(p.ContactInfo.ContactName, t.client.config.PlatformPublicKey)
	if err != nil {
		return nil, err
	}

	if p.AccountInfo != nil {
		p.AccountInfo.AccountNumber, err = tool.Encrypt(p.AccountInfo.AccountNumber, t.client.config.PlatformPublicKey)
		if err != nil {
			return nil, err
		}
		p.AccountInfo.AccountName, err = tool.Encrypt(p.AccountInfo.AccountName, t.client.config.PlatformPublicKey)
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
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
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
func (t *Ecommerce) ApplyQuery(p *EcommerceApplyQuery) (*EcommerceApplyQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/applyments/" + p.ApplymentID
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
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
func (t *Ecommerce) ModifySettlement(p *EcommerceModifySettlement) error {
	// 加密银行卡号
	if p.AccountNumber != "" {
		AccountNumberMD, err := tool.Encrypt(p.AccountNumber, t.client.config.PlatformPublicKey)
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
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return err
		}
	}

	return nil
}

// QuerySettlement 查询结算信息
func (t *Ecommerce) QuerySettlement(p *EcommerceQuerySettlement) (*EcommerceQuerySettlementResp, error) {

	// 发起请求
	urlPath := "/v3/apply4sub/sub_merchants/" + p.SubMchid + "/settlement"
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output EcommerceQuerySettlementResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
