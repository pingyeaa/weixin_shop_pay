package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
)

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "\n-----END RSA PRIVATE KEY-----"
)

// Signature 签名
func Signature(urlPath string, requestBody string, privateKey string, mchid string, serialNo string) (string, error) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return "", err
	}
	method := "POST"
	timestamp := time.Now().UnixNano()
	nonceStr := node.Generate()
	signString := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, urlPath, timestamp, nonceStr, requestBody)
	log.Println(signString)
	sign := RsaSign(signString, privateKey, crypto.SHA256)
	return fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", mchid, nonceStr, sign, timestamp, serialNo), nil
}

func RsaSign(signContent string, privateKey string, hash crypto.Hash) string {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}
