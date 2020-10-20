package weixin_shop_pay

// Config 配置
type Config struct {
	SpAppID  string // 服务商公众号ID
	SpMchID  string // 服务商户号
	KeyPath  string // 私钥地址
	SerialNo string // 证书序列号
}

// NewClient 创建客户端
func NewClient(c *Config) *Client {
	return &Client{
		c,
	}
}

// Client 客户端
type Client struct {
	Config *Config
}
