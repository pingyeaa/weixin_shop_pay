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
func (c *Ecommerce) Apply(p *params.EcommerceApply) (*params.EcommerceApplyResp, error) {

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
	var output params.EcommerceApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ApplyQuery 二级商户进件查询
func (c *Ecommerce) ApplyQuery(p *params.EcommerceApplyQuery) (*params.EcommerceApplyQueryResp, error) {

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
	var output params.EcommerceApplyQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
