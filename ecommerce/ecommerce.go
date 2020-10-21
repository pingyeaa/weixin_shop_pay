package ecommerce

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pingyeaa/weixin-shop-pay/config"

	"github.com/pingyeaa/weixin-shop-pay/params"
	"github.com/pingyeaa/weixin-shop-pay/tools"
)

// Ecommerce 二级商户进件
type Ecommerce struct {
	Config *config.Config
}

// Apply 二级商户进件
func (c *Ecommerce) Apply(p *params.EcommerceApply) (*ApplyResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/applyments"
	resp, err := tools.PostRequest(c.Config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output ApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyResp 二级商户进件返回参数
type ApplyResp struct {
	ApplymentID  string `json:"applyment_id"`   // 微信支付申请单号
	OutRequestNo string `json:"out_request_no"` // 业务申请编号
}

// ApplyQuery 二级商户进件查询
func (c *Ecommerce) ApplyQuery(p *params.EcommerceApplyQuery) (*ApplyQueryResp, error) {

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
	signature, err := tools.Signature(urlPath, string(dataJsonByte), string(keyByte), c.Config.SpMchID, c.Config.SerialNo)
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
	var output ApplyQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyQueryResp 二级商户进件查询
type ApplyQueryResp struct {
	ApplymentState     string `json:"applyment_state"`      // 申请状态
	ApplymentStateDesc string `json:"applyment_state_desc"` // 申请状态描述
	SignURL            string `json:"sign_url"`             // 签约链接
	SubMchid           string `json:"sub_mchid"`            // 电商平台二级商户号
	AccountValidation  struct {
		AccountName              int64  `json:"account_name"`               // 付款户名
		AccountNo                string `json:"account_no"`                 // 付款卡号
		PayAmount                string `json:"pay_amount"`                 // 汇款金额
		DestinationAccountNumber string `json:"destination_account_number"` // 收款卡号
		DestinationAccountName   string `json:"destination_account_name"`   // 收款户名
		DestinationAccountBank   string `json:"destination_account_bank"`   // 开户银行
		City                     string `json:"city"`                       // 省市信息
		Remark                   string `json:"remark"`                     // 备注信息
		Deadline                 string `json:"deadline"`                   // 汇款截止日期
	} `json:"account_validation"` // 汇款账户验证信息
	AuditDetail []struct {
		ParamName    string `json:"param_name"`    // 参数名称
		RejectReason string `json:"reject_reason"` // 驳回原因
	} `json:"audit_detail"` // 驳回原因详情
	LegalValidationURL string `json:"legal_validation_url"` // 法人验证链接
	OutRequestNo       string `json:"out_request_no"`       // 业务申请编号
	ApplymentID        int64  `json:"applyment_id"`         // 微信支付申请单号
}
