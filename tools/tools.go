package tools

import (
	"crypto/sha256"
	"encoding/hex"
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
