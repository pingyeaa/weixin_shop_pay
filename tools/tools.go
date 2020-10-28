package tools

import (
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
	"log"
	"os"
)

// GetFileHash 文件sha256哈希
func GetFileHash(path string) string {
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
func Encrypt(content string, rsaPublicKey string) (string, error) {
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

func AesDecrypt(content string, key string, nonce string, associatedData string) ([]byte, error) {
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
