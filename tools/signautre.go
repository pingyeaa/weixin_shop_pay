package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/snowflake"
)

// Signature 签名
func Signature(urlPath string, requestBody string, privateKey string, mchid string, serialNo string) (string, error) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return "", err
	}
	method := "POST"
	timestamp := time.Now().Unix()
	nonceStr := node.Generate()
	signString := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, urlPath, timestamp, nonceStr, requestBody)
	log.Println("签名原文", signString)
	sign, err := RsaSignWithSha256([]byte(signString), privateKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", mchid, nonceStr, sign, timestamp, serialNo), nil
}

//签名
func RsaSignWithSha256(data []byte, privateKey string) (string, error) {
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
