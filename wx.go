package weixin_shop_pay

import (
	"weixin_shop_pay/balance"
	"weixin_shop_pay/ecommerce"
	"weixin_shop_pay/normal_pay"
	"weixin_shop_pay/profitsharing"
	"weixin_shop_pay/refund"
	"weixin_shop_pay/withdraw"
)

// Config 配置
type Config struct {
	SpAppID  string // 服务商公众号ID
	SpMchID  string // 服务商户号
	KeyPath  string // 私钥地址
	SerialNo string // 证书序列号
}

// Domain 请求域名
var Domain = "https://api.mch.weixin.qq.com/"

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

// NormalPay 普通支付
func (c *Client) NormalPay() *normal_pay.NormalPay {
	return &normal_pay.NormalPay{Config: c.Config}
}

// Ecommerce 二级商户进件
func (c *Client) Ecommerce() *ecommerce.Ecommerce {
	return &ecommerce.Ecommerce{Config: c.Config}
}

// ProfitSharing 分账
func (c *Client) ProfitSharing() *profitsharing.ProfitSharing {
	return &profitsharing.ProfitSharing{Config: c.Config}
}

// Refund 退款
func (c *Client) Refund() *refund.Refund {
	return &refund.Refund{Config: c.Config}
}

// Balance 余额
func (c *Client) Balance() *balance.Balance {
	return &balance.Balance{Config: c.Config}
}

// Withdraw 提现
func (c *Client) Withdraw() *withdraw.Withdraw {
	return &withdraw.Withdraw{Config: c.Config}
}
