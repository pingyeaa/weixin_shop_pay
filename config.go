package weixin_shop_pay

// Config 配置
type Config struct {
	SpAppID           string // 服务商公众号ID
	SpMchID           string // 服务商户号
	KeyPath           string // 私钥地址
	SerialNo          string // 商户证书序列号
	SecretKey         string // APIV3密钥
	PlatformPublicKey string // 平台公钥
	PlatformSerialNo  string // 平台证书序列号
}

// Domain 请求域名
var Domain = "https://api.mch.weixin.qq.com"
