package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"

	config2 "github.com/pingyeaa/weixin_shop_pay/config"
)

// PostRequest 请求接口
func PostRequest(config *config2.Config, urlPath string, dataJsonByte []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", config2.Domain+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := Signature("POST", urlPath, string(dataJsonByte), string(keyByte), config.SpMchID, config.SerialNo)
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
func GetRequest(config *config2.Config, urlPath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", config2.Domain+urlPath, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := Signature("GET", urlPath, "", string(keyByte), config.SpMchID, config.SerialNo)
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
