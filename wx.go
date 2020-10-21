package weixin_shop_pay

import (
	"github.com/pingyeaa/weixin-shop-pay/balance"
	"github.com/pingyeaa/weixin-shop-pay/config"
	"github.com/pingyeaa/weixin-shop-pay/ecommerce"
	"github.com/pingyeaa/weixin-shop-pay/pay"
	"github.com/pingyeaa/weixin-shop-pay/profitsharing"
	"github.com/pingyeaa/weixin-shop-pay/refund"
	"github.com/pingyeaa/weixin-shop-pay/withdraw"
)

// NewClient 创建客户端
func NewClient(c *config.Config) *Client {
	return &Client{
		c,
	}
}

// NewConfig 实例化配置
func NewConfig() *config.Config {
	return &config.Config{}
}

// Client 客户端
type Client struct {
	Config *config.Config
}

// Pay 普通支付
func (c *Client) Pay() *pay.Pay {
	return &pay.Pay{Config: c.Config}
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
