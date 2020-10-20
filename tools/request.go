package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"weixin_shop_pay"
)

// 请求接口
func Request(urlPath string, dataJsonByte []byte, keyPath string) (*http.Response, error) {
	req, err := http.NewRequest("POST", weixin_shop_pay.Domain+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(keyPath)
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
	return resp, nil
}
