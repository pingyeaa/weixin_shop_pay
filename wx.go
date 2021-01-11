package weixin_shop_pay

// NewClient 创建客户端
func NewClient(c *Config) *Client {
	return &Client{
		c,
	}
}

// NewConfig 实例化配置
func NewConfig() *Config {
	return &Config{}
}

// Client 客户端
type Client struct {
	Config *Config
}

// Pay 普通支付
func (c *Client) Pay() *Pay {
	return &Pay{Config: c.Config}
}

// Ecommerce 二级商户进件
func (c *Client) Ecommerce() *Ecommerce {
	return &Ecommerce{Config: c.Config}
}

// ProfitSharing 分账
func (c *Client) ProfitSharing() *ProfitSharing {
	return &ProfitSharing{Config: c.Config}
}

// Refund 退款
func (c *Client) Refund() *Refund {
	return &Refund{Config: c.Config}
}

// Balance 余额
func (c *Client) Balance() *Balance {
	return &Balance{Config: c.Config}
}

// Withdraw 提现
func (c *Client) Withdraw() *Withdraw {
	return &Withdraw{Config: c.Config}
}

// Common 公共接口
func (c *Client) Common() *Common {
	return &Common{Config: c.Config}
}

// Cert 证书
func (c *Client) Cert() *Cert {
	return &Cert{Config: c.Config}
}
