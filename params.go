package weixin_shop_pay

// OrderParams 下单参数
type OrderParams struct {
	SpAppID     string `json:"sp_appid"`     // 服务商公众号ID
	SpMchID     string `json:"sp_mchid"`     // 服务商户号
	SubAppID    string `json:"sub_appid"`    // 二级商户公众号ID
	SubMchID    string `json:"sub_mchid"`    // 二级商户号
	Description string `json:"description"`  // 商品描述
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号
	TimeExpire  string `json:"time_expire"`  // 交易结束时间
	Attach      string `json:"attach"`       // 附加数据
	NotifyURL   string `json:"notify_url"`   // 通知地址
	GoodsTag    string `json:"goods_tag"`    // 订单优惠标记
	SettleInfo  struct {
		ProfitSharing bool  `json:"profit_sharing"` // 是否制定分账
		SubsidyAmount int64 `json:"subsidy_amount"` // 补差金额
	} `json:"settle_info"` // 结算信息
	Amount    *Amount    `json:"amount"`     // 订单金额
	Payer     *Payer     `json:"payer"`      // 支付者
	Detail    *Detail    `json:"detail"`     // 优惠功能
	SceneInfo *SceneInfo `json:"scene_info"` // 场景信息
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

// QueryOrderParams 查询订单参数
type QueryOrderParams struct {
	SpMchID       string `json:"sp_mchid"`       // 服务商户号
	SubAppID      string `json:"sub_appid"`      // 二级商户公众号ID
	TransactionID string `json:"transaction_id"` // 订单号
}

// EcommerceApplyParams 二级商户进件参数
type EcommerceApplyParams struct {
	OutRequestNo        string `json:"out_request_no"`    // 业务申请编号
	OrganizationType    string `json:"organization_type"` // 主体类型
	BusinessLicenseInfo struct {
		BusinessLicenseCopy   string `json:"business_license_copy"`   // 证件扫描件
		BusinessLicenseNumber string `json:"business_license_number"` // 证件注册号
		MerchantName          string `json:"merchant_name"`           // 商户名称
		LegalPerson           string `json:"legal_person"`            // 经营者/法定代表人姓名
		CompanyAddress        string `json:"company_address	"`        // 注册地址
		BusinessTime          string `json:"business_time"`           // 营业期限
	} `json:"business_license_info"` // 营业执照/登记证书信息
	OrganizationCertInfo struct {
		OrganizationCopy   string `json:"organization_copy"`   // 组织机构代码证照片
		OrganizationNumber string `json:"organization_number"` // 组织机构代码
		OrganizationTime   string `json:"organization_time"`   // 组织机构代码有效期限
	} `json:"organization_cert_info"` // 组织机构代码证信息
	IDDocType  string `json:"id_doc_type"` // 经营者/法人证件类型
	IDCardInfo struct {
		IDCardCopy      string `json:"id_card_copy"`       // 身份证人像面照片
		IDCardNational  string `json:"id_card_national"`   // 身份证国徽面照片
		IDCardName      string `json:"id_card_name"`       // 身份证姓名
		IDCardNumber    string `json:"id_card_number"`     // 身份证号码
		IDCardValidTime string `json:"id_card_valid_time"` // 身份证有效期限
	} `json:"id_card_info"` // 经营者/法人身份证信息
	IDDocInfo struct {
		IDDocName    string `json:"id_doc_name"`    // 证件姓名
		IDDocNumber  string `json:"id_doc_number"`  // 证件号码
		IDDocCopy    string `json:"id_doc_copy"`    // 证件照片
		DocPeriodEnd string `json:"doc_period_end"` // 证件结束日期
	} `json:"id_doc_info"` // 经营者/法人其他类型证件信息
	NeedAccountInfo bool `json:"need_account_info"` // 是否填写结算银行账户
	AccountInfo     struct {
		BankAccountType string `json:"bank_account_type"` // 账户类型
		AccountBank     string `json:"account_bank"`      // 开户银行
		AccountName     string `json:"account_name"`      // 开户名称
		BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
		BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
		BankName        string `json:"bank_name"`         // 开户银行全称
		AccountNumber   string `json:"account_number"`    // 银行账号
	} `json:"account_info"` // 结算银行账户
	ContactInfo struct {
		ContactType         string `json:"contact_type"`           // 超级管理员类型
		ContactName         string `json:"contact_name"`           // 超级管理员姓名
		ContactIDCardNumber string `json:"contact_id_card_number"` // 超级管理员身份证件号码
		ContactPhone        string `json:"mobile_phone"`           // 超级管理员手机
		ContactEmail        string `json:"contact_email"`          // 超级管理员邮箱
	} `json:"contact_info"` // 超级管理员
	SalesSceneInfo struct {
		StoreName           string `json:"store_name"`             // 店铺名称
		StoreURL            string `json:"store_url"`              // 店铺链接
		StoreQrCode         string `json:"store_qr_code"`          // 店铺二维码
		MiniProgramSubAPPID string `json:"mini_program_sub_appid"` // 小程序APPID
	} `json:"sales_scene_info"` // 店铺信息
	// merchant_shortname
	MerchantShortname    string `json:"merchant_shortname"`     // 商户简称
	Qualifications       string `json:"qualifications"`         // 特殊资质
	BusinessAdditionPics string `json:"business_addition_pics"` // 补充材料
	BusinessAdditionDesc string `json:"business_addition_desc"` // 补充说明
}

// EcommerceApplyQueryParams 二级商户进件查询参数
type EcommerceApplyQueryParams struct {
	ApplymentID string `json:"applyment_id"` // 微信支付申请单号
}

// ReceiversAddParams 添加分账接收方参数
type ReceiversAddParams struct {
	Appid         string `json:"appid"`          // 公众账号ID
	Type          string `json:"type"`           // 接收方类型
	Account       string `json:"account"`        // 接收方账号
	Name          string `json:"name"`           // 接收方名称
	EncryptedName string `json:"encrypted_name"` // 接收方名称的密文
	RelationType  string `json:"relation_type"`  // 与分账方的关系类型
}

// ProfitSharingOrders 请求分账参数
type ProfitSharingOrders struct {
	Appid         string                          `json:"appid"`          // 公众账号ID
	SubMchid      string                          `json:"sub_mchid"`      // 二级商户号
	TransactionID string                          `json:"transaction_id"` // 微信订单号
	OutTradeNo    string                          `json:"out_trade_no"`   // 商户分账单号
	Receivers     []*ProfitSharingOrdersReceivers `json:"receivers"`      // 分账接收方列表
	Finish        bool                            `json:"finish"`         // 是否分账完成
}

// ProfitSharingOrdersReceivers 请求分账-分账接收方列表
type ProfitSharingOrdersReceivers struct {
	Type            string `json:"type"`             // 分账接收方类型
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号
	Amount          int    `json:"amount"`           // 分账金额
	Description     string `json:"description"`      // 分账描述
	ReceiverName    string `json:"receiver_name"`    // 分账姓名
}

// ProfitSharingQuery 分账-查询分账结果
type ProfitSharingQuery struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
}

// ProfitSharingFinishOrder 分账-完结分账
type ProfitSharingFinishOrder struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	Description   string `json:"description"`    // 分账描述
}

// RefundApply 申请退款
type RefundApply struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	SpAppid       string `json:"sp_appid"`       // 电商平台APPID
	SubAppid      string `json:"sub_appid"`      // 二级商户APPID
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户订单号
	OutRefundNo   string `json:"out_refund_no"`  // 商户退款单号
	Reason        string `json:"reason"`         // 退款原因
	Amount        struct {
		Refund   int    `json:"refund"`   // 退款金额
		Total    int    `json:"total"`    // 原订单金额
		Currency string `json:"currency"` // 退款币种
	} `json:"amount"` // 订单金额
	NotifyURL string `json:"notify_url"` // 退款结果回调URL
}

// RefundQuery 退款查询
type RefundQuery struct {
	RefundID string `json:"refund_id"` // 微信退款单号
	SubMchid string `json:"sub_mchid"` // 二级商户号
}
