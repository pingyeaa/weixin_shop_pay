package weixin_shop_pay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// Common 通用接口
type Common struct {
	client *Client
}

// ImageUpload 图片上传
func (t *Common) ImageUpload(p *CommonImageUpload) (*CommonImageUploadResp, error) {
	var res CommonImageUploadResp

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(t.client.config.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("私钥文件读取失败：%s", err)
	}

	// 读取图片内容
	imageFile, err := os.Open(p.FilePath)
	if err != nil {
		return nil, err
	}
	fileByte, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	// 计算文件哈希256
	dataJsonByte, err := json.Marshal(map[string]string{
		"filename": imageFile.Name(),
		"sha256":   tool.GetFileHash(p.FilePath),
	})

	// 签名
	urlPath := "/v3/merchant/media/upload"
	signature, err := tool.Signature("POST", urlPath, string(dataJsonByte), string(keyByte), t.client.config.SpMchID, t.client.config.SerialNo)
	if err != nil {
		return nil, err
	}

	// 拼接请求体
	imageFileExt := path.Ext(p.FilePath)
	requestBody := []byte(fmt.Sprintf("--boundary\r\n"+
		"Content-Disposition: form-data; name=\"meta\";\r\n"+
		"Content-Type: application/json\r\n"+
		"\r\n"+
		"%s\r\n"+
		"--boundary\r\n"+
		"Content-Disposition: form-data; name=\"file\"; filename=\"%s\";\r\n"+
		"Content-Type: image/%s\r\n"+
		"\r\n"+
		"%s\r\n"+
		"--boundary--", string(dataJsonByte), imageFile.Name(), strings.Replace(imageFileExt, ".", "", -1), fileByte))

	// 设置请求头
	req, err := http.NewRequest("POST", Domain+urlPath, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data;boundary=boundary")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "WECHATPAY2-SHA256-RSA2048 "+signature)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &res, nil
	}
	log.Println("上传图片接口响应信息", string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		return &res, nil
	}

	return &res, nil
}
