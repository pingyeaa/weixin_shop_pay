package weixin_shop_pay

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
)

var tool *Tool

// Tool 工具
type Tool struct{}

// PostRequest 请求接口
func (t *Tool) PostRequest(config *Config, urlPath string, dataJsonByte []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", Domain+urlPath, bytes.NewBuffer(dataJsonByte))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := t.Signature("POST", urlPath, string(dataJsonByte), string(keyByte), config.SpMchID, config.SerialNo)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "WECHATPAY2-SHA256-RSA2048 "+signature)
	req.Header.Set("Wechatpay-Serial", config.PlatformSerialNo)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("POST请求报错", err.Error())
		return nil, err
	}
	return resp, nil
}

// GetRequest .
func (t *Tool) GetRequest(config *Config, urlPath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", Domain+urlPath, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	// 读取私钥文件
	keyByte, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 签名
	signature, err := t.Signature("GET", urlPath, "", string(keyByte), config.SpMchID, config.SerialNo)
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

// Signature 签名
func (t *Tool) Signature(method string, urlPath string, requestBody string, privateKey string, mchid string, serialNo string) (string, error) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	nonceStr := node.Generate()
	signString := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, urlPath, timestamp, nonceStr, requestBody)
	log.Println("签名原文", signString)
	sign, err := t.RsaSignWithSha256([]byte(signString), privateKey)
	if err != nil {

		return "", err
	}
	return fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", mchid, nonceStr, sign, timestamp, serialNo), nil
}

//签名
func (t *Tool) RsaSignWithSha256(data []byte, privateKey string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write(data)
	hash := h.Sum(nil)
	log.Println("哈希", hash)
	signature, err := rsa.SignPKCS1v15(rand.Reader, pri.(*rsa.PrivateKey), crypto.SHA256, hash)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// GetFileHash 文件sha256哈希
func (t *Tool) GetFileHash(path string) string {
	//创建一个可操作的sha256对象
	hash := sha256.New()
	//打开所需校验的文件
	fp, _ := os.Open(path)
	defer fp.Close()
	buf := make([]byte, 1024)
	//持续读取文件内容
	for {
		n, _ := fp.Read(buf)
		if n == 0 {
			break
		}
		//将每次读取到的数据都添加到hash中去
		hash.Write(buf[:n])
	}
	//最后来一次大汇总
	result := hash.Sum(nil)
	//转化为十六进制后输出到屏幕
	return hex.EncodeToString(result)
}

// Encrypt 参数值加密
func (t *Tool) Encrypt(content string, rsaPublicKey string) (string, error) {
	secretMessage := []byte(content)
	rng := rand.Reader

	block, _ := pem.Decode([]byte(rsaPublicKey))
	if block == nil {
		return "", errors.New("failed to decode PEM block containing public key")
	}

	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	pub := cert.PublicKey.(*rsa.PublicKey)

	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, pub, secretMessage, nil)
	if err != nil {
		return "", fmt.Errorf("敏感信息加密失败：%s", err)
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	return ciphertext, nil
}

func (t *Tool) AesDecrypt(content string, key string, nonce string, associatedData string) ([]byte, error) {
	log.Println(content, string(key), string(nonce), string(associatedData))
	plaintext, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext, err := aead.Open(nil, []byte(nonce), plaintext, []byte(associatedData))
	if err != nil {
		panic(err.Error())
	}
	return ciphertext, nil
}
