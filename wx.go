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

// Amount 订单金额
type Amount struct {
	Total    int    `json:"total"`    // 订单总金额
	Currency string `json:"currency"` // 货币类型
}

// Payer 支付者
type Payer struct {
	SpOpenid  string `json:"sp_openid"`  // 用户服务标识
	SubOpenid string `json:"sub_openid"` // 用户子标识
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int    `json:"cost_price"` // 订单原价
	InvoiceID   string `json:"invoice_id"` // 商家小票
	GoodsDetail []struct {
		MerchantsGoodsID string `json:"merchants_goods_id"` // 商户侧商品编码
		WechatpayGoodsID string `json:"wechatpay_goods_id"` // 微信侧商品编码
		GoodsName        string `json:"goods_name"`         // 商品的实际名称
		Quantity         int    `json:"quantity"`           // 商品数量
		UnitPrice        int    `json:"unit_price"`         // 商品单价
	} `json:"goods_detail"` // 单品列表
}

// SceneInfo 场景信息
type SceneInfo struct {
	PayerClientIp string `json:"payer_client_ip"` // 用户终端IP
	DeviceID      string `json:"device_id"`       // 商户端设备号
	StoreInfo     struct {
		ID       string `json:"id"`        // 门店编号
		Name     string `json:"name"`      // 门店名称
		AreaCode string `json:"area_code"` // 地区编码
		Address  string `json:"address"`   // 详细地址
	} `json:"store_info"` // 商户门店信息
}
