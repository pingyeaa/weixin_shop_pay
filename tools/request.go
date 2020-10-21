package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"

	weixin_shop_pay "github.com/pingyeaa/weixin-shop-pay"
)

// PostRequest 请求接口
func PostRequest(config *weixin_shop_pay.Config, urlPath string, dataJsonByte []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", weixin_shop_pay.Domain+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := Signature(urlPath, string(dataJsonByte), string(keyByte), config.SpMchID, config.SerialNo)
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
	return resp, nil
}

// GetRequest .
func GetRequest(config *weixin_shop_pay.Config, urlPath string, dataJsonByte []byte) (*http.Response, error) {
	req, err := http.NewRequest("GET", weixin_shop_pay.Domain+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := Signature(urlPath, string(dataJsonByte), string(keyByte), config.SpMchID, config.SerialNo)
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
	return resp, nil
}
