package weixin_shop_pay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// EcommerceApplyParams 二级商户进件参数
type EcommerceApplyParams struct {
	OutRequestNo        string `json:"out_request_no"`    // 业务申请编号
	OrganizationType    string `json:"organization_type"` // 主体类型
	BusinessLicenseInfo struct {
		BusinessLicenseCopy   string `json:"business_license_copy"`   // 证件扫描件
		BusinessLicenseNumber string `json:"business_license_number"` // 证件注册号
		MerchantName          string `json:"merchant_name"`           // 商户名称
		LegalPerson           string `json:"legal_person"`            // 经营者/法定代表人姓名
		CompanyAddress        string `json:"company_address	"`        // 注册地址
		BusinessTime          string `json:"business_time"`           // 营业期限
	} `json:"business_license_info"` // 营业执照/登记证书信息
	OrganizationCertInfo struct {
		OrganizationCopy   string `json:"organization_copy"`   // 组织机构代码证照片
		OrganizationNumber string `json:"organization_number"` // 组织机构代码
		OrganizationTime   string `json:"organization_time"`   // 组织机构代码有效期限
	} `json:"organization_cert_info"` // 组织机构代码证信息
	IDDocType  string `json:"id_doc_type"` // 经营者/法人证件类型
	IDCardInfo struct {
		IDCardCopy      string `json:"id_card_copy"`       // 身份证人像面照片
		IDCardNational  string `json:"id_card_national"`   // 身份证国徽面照片
		IDCardName      string `json:"id_card_name"`       // 身份证姓名
		IDCardNumber    string `json:"id_card_number"`     // 身份证号码
		IDCardValidTime string `json:"id_card_valid_time"` // 身份证有效期限
	} `json:"id_card_info"` // 经营者/法人身份证信息
	IDDocInfo struct {
		IDDocName    string `json:"id_doc_name"`    // 证件姓名
		IDDocNumber  string `json:"id_doc_number"`  // 证件号码
		IDDocCopy    string `json:"id_doc_copy"`    // 证件照片
		DocPeriodEnd string `json:"doc_period_end"` // 证件结束日期
	} `json:"id_doc_info"` // 经营者/法人其他类型证件信息
	NeedAccountInfo bool `json:"need_account_info"` // 是否填写结算银行账户
	AccountInfo     struct {
		BankAccountType string `json:"bank_account_type"` // 账户类型
		AccountBank     string `json:"account_bank"`      // 开户银行
		AccountName     string `json:"account_name"`      // 开户名称
		BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
		BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
		BankName        string `json:"bank_name"`         // 开户银行全称
		AccountNumber   string `json:"account_number"`    // 银行账号
	} `json:"account_info"` // 结算银行账户
	ContactInfo struct {
		ContactType         string `json:"contact_type"`           // 超级管理员类型
		ContactName         string `json:"contact_name"`           // 超级管理员姓名
		ContactIDCardNumber string `json:"contact_id_card_number"` // 超级管理员身份证件号码
		ContactPhone        string `json:"mobile_phone"`           // 超级管理员手机
		ContactEmail        string `json:"contact_email"`          // 超级管理员邮箱
	} `json:"contact_info"` // 超级管理员
	SalesSceneInfo struct {
		StoreName           string `json:"store_name"`             // 店铺名称
		StoreURL            string `json:"store_url"`              // 店铺链接
		StoreQrCode         string `json:"store_qr_code"`          // 店铺二维码
		MiniProgramSubAPPID string `json:"mini_program_sub_appid"` // 小程序APPID
	} `json:"sales_scene_info"` // 店铺信息
	// merchant_shortname
	MerchantShortname    string `json:"merchant_shortname"`     // 商户简称
	Qualifications       string `json:"qualifications"`         // 特殊资质
	BusinessAdditionPics string `json:"business_addition_pics"` // 补充材料
	BusinessAdditionDesc string `json:"business_addition_desc"` // 补充说明
}

// EcommerceApply 二级商户进件
func (c *Client) EcommerceApply(p *EcommerceApplyParams) (*EcommerceApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/applyments"
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/"+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := Signature(urlPath, string(dataJsonByte), string(keyByte), c.Config.SpMchID, c.Config.SerialNo)
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
	var output EcommerceApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// EcommerceApplyResp 二级商户进件返回参数
type EcommerceApplyResp struct {
	ApplymentID  string `json:"applyment_id"`   // 微信支付申请单号
	OutRequestNo string `json:"out_request_no"` // 业务申请编号
}
