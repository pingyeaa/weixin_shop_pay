package weixin_shop_pay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Cert 证书
type Cert struct {
	Config *Config
}

// Certificates 平台证书列表
func (c *Cert) Certificates() (*CertCertificatesResp, error) {

	// 发起请求
	urlPath := "/v3/certificates"
	resp, err := tool.GetRequest(c.Config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(respData))
	var cipherResp CipherResp
	err = json.Unmarshal(respData, &cipherResp)
	if err != nil {
		return nil, err
	}

	output := CertCertificatesResp{}
	for _, data := range cipherResp.Data {
		log.Println(data.EncryptCertificate.Ciphertext)
		log.Println(c.Config.SecretKey)
		decryptContent, err := tool.AesDecrypt(data.EncryptCertificate.Ciphertext, c.Config.SecretKey, data.EncryptCertificate.Nonce, data.EncryptCertificate.AssociatedData)
		if err != nil {
			return nil, fmt.Errorf("证书结果解密失败：%s", err)
		}
		log.Println("解密数据", string(decryptContent))
		output.List = append(output.List, CertCertificatesListResp{
			SerialNo:  data.SerialNo,
			PublicKey: string(decryptContent),
		})
	}

	return &output, nil
}
