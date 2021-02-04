package weixin_shop_pay

import "encoding/json"

// NewClient 创建客户端
func NewClient(config *Config) *Client {
	return &Client{
		config:        config,
		errorResponse: nil,
	}
}

// NewConfig 实例化配置
func NewConfig() *Config {
	return &Config{}
}

// Client 客户端
type Client struct {
	config        *Config
	errorResponse *errorResponse
}

// HTTP CODE不等于200，或204时的错误返回参数
type errorResponse struct {
	Code    string
	Message string
	Detail  struct {
		Field    string
		Value    string
		Issue    string
		Location string
	}
}

// Pay 普通支付
func (t *Client) Pay() *Pay {
	return &Pay{client: t}
}

// Ecommerce 二级商户进件
func (t *Client) Ecommerce() *Ecommerce {
	return &Ecommerce{client: t}
}

// ProfitSharing 分账
func (t *Client) ProfitSharing() *ProfitSharing {
	return &ProfitSharing{client: t}
}

// Refund 退款
func (t *Client) Refund() *Refund {
	return &Refund{client: t}
}

// Balance 余额
func (t *Client) Balance() *Balance {
	return &Balance{client: t}
}

// Withdraw 提现
func (t *Client) Withdraw() *Withdraw {
	return &Withdraw{client: t}
}

// Common 公共接口
func (t *Client) Common() *Common {
	return &Common{client: t}
}

// Cert 证书
func (t *Client) Cert() *Cert {
	return &Cert{client: t}
}

// setErrorResponse 设置错误响应信息
func (t *Client) setErrorResponse(resp []byte) error {
	var errorResponse *errorResponse
	err := json.Unmarshal(resp, errorResponse)
	if err != nil {
		return err
	}
	t.errorResponse = errorResponse
	return nil
}

// GetErrorResponse 获取微信错误响应信息
func (t *Client) GetErrorResponse() *errorResponse {
	return t.errorResponse
}
