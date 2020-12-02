package params

// PayOrder 下单
type PayOrder struct {
	SpAppID     string              `json:"sp_appid"`     // 服务商公众号ID
	SpMchID     string              `json:"sp_mchid"`     // 服务商户号
	SubAppID    string              `json:"sub_appid"`    // 二级商户公众号ID
	SubMchID    string              `json:"sub_mchid"`    // 二级商户号
	Description string              `json:"description"`  // 商品描述
	OutTradeNo  string              `json:"out_trade_no"` // 商户订单号
	TimeExpire  string              `json:"time_expire"`  // 交易结束时间
	Attach      string              `json:"attach"`       // 附加数据
	NotifyURL   string              `json:"notify_url"`   // 通知地址
	GoodsTag    string              `json:"goods_tag"`    // 订单优惠标记
	SettleInfo  *PayOrderSettleInfo `json:"settle_info"`  // 结算信息
	Amount      *PayOrderAmount     `json:"amount"`       // 订单金额
	Payer       *PayOrderPayer      `json:"payer"`        // 支付者
	//Detail      *PayOrderDetail     `json:"detail"`       // 优惠功能
	//SceneInfo   *PayOrderSceneInfo  `json:"scene_info"`   // 场景信息
}

// PayOrderSettleInfo 下单-结算信息
type PayOrderSettleInfo struct {
	ProfitSharing bool  `json:"profit_sharing"` // 是否制定分账
	SubsidyAmount int64 `json:"subsidy_amount"` // 补差金额
}

// PayOrderAmount 下单-订单金额
type PayOrderAmount struct {
	Total    int    `json:"total"`    // 订单总金额
	Currency string `json:"currency"` // 货币类型
}

// PayOrderPayer 下单-支付者
type PayOrderPayer struct {
	SpOpenid  string `json:"sp_openid"`  // 用户服务标识
	SubOpenid string `json:"sub_openid"` // 用户子标识
}

// PayOrderDetail 下单-优惠功能
type PayOrderDetail struct {
	//CostPrice   int                          `json:"cost_price"`   // 订单原价
	//InvoiceID   string                       `json:"invoice_id"`   // 商家小票
	GoodsDetail []*PayOrderDetailGoodsDetail `json:"goods_detail"` // 单品列表
}

// PayOrderDetailGoodsDetail 下单-优惠功能-单品列表
type PayOrderDetailGoodsDetail struct {
	MerchantGoodsID string `json:"merchant_goods_id"` // 商户侧商品编码
	//WechatpayGoodsID string `json:"wechatpay_goods_id"` // 微信侧商品编码
	//GoodsName        string `json:"goods_name"`         // 商品的实际名称
	Quantity  int `json:"quantity"`   // 商品数量
	UnitPrice int `json:"unit_price"` // 商品单价
}

// PayOrderSceneInfo 下单-场景信息
type PayOrderSceneInfo struct {
	PayerClientIp string                      `json:"payer_client_ip"` // 用户终端IP
	DeviceID      string                      `json:"device_id"`       // 商户端设备号
	StoreInfo     *PayOrderSceneInfoStoreInfo `json:"store_info"`      // 商户门店信息
}

// PayOrderSceneInfoStoreInfo 下单-场景信息-商户门店信息
type PayOrderSceneInfoStoreInfo struct {
	ID       string `json:"id"`        // 门店编号
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 地区编码
	Address  string `json:"address"`   // 详细地址
}

// PayOrderResp 下单返回数据
type PayOrderResp struct {
	PrepayId string `json:"prepay_id"`
}

// PayQueryOrderTransaction 微信订单查询
type PayQueryOrderTransaction struct {
	SpMchID       string `json:"sp_mchid"`       // 服务商户号
	SubMchID      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 订单号
}

// PayQueryOrderOutTradeNo 商户订单查询
type PayQueryOrderOutTradeNo struct {
	SpMchID    string `json:"sp_mchid"`     // 服务商户号
	SubMchID   string `json:"sub_mchid"`    // 二级商户号
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
}

// PayQueryOrderResp 微信订单查询返回参数
type PayQueryOrderResp struct {
	SpAppID        string             `json:"sp_appid"`         // 服务商公众号ID
	SpMchID        string             `json:"sp_mchid"`         // 服务商户号
	SubAppID       string             `json:"sub_appid"`        // 二级商户公众号ID
	SubMchID       string             `json:"sub_mchid"`        // 二级商户号
	OutTradeNo     string             `json:"out_trade_no"`     // 商户订单号
	TransactionID  string             `json:"transaction_id"`   // 微信支付订单号
	TradeType      string             `json:"trade_type"`       // 交易类型
	TradeState     string             `json:"trade_state"`      // 交易状态
	TradeStateDesc string             `json:"trade_state_desc"` // 交易状态描述
	BankType       string             `json:"bank_type"`        // 付款银行
	Attach         string             `json:"attach"`           // 附加数据
	SuccessTime    string             `json:"success_time"`     // 付款完成时间
	Payer          *PayOrderPayer     `json:"payer"`            // 支付者
	Amount         *PayOrderAmount    `json:"amount"`           // 订单金额
	Detail         *PayOrderDetail    `json:"detail"`           // 优惠功能
	SceneInfo      *PayOrderSceneInfo `json:"scene_info"`       // 场景信息
}

// EcommerceApply 二级商户进件
type EcommerceApply struct {
	OutRequestNo         string                              `json:"out_request_no"`         // 业务申请编号
	OrganizationType     string                              `json:"organization_type"`      // 主体类型
	BusinessLicenseInfo  *EcommerceApplyBusinessLicenseInfo  `json:"business_license_info"`  // 营业执照/登记证书信息
	OrganizationCertInfo *EcommerceApplyOrganizationCertInfo `json:"organization_cert_info"` // 组织机构代码证信息
	IDDocType            string                              `json:"id_doc_type"`            // 经营者/法人证件类型
	IDCardInfo           *EcommerceApplyIDCardInfo           `json:"id_card_info"`           // 经营者/法人身份证信息
	IDDocInfo            *EcommerceApplyIDDocInfo            `json:"id_doc_info"`            // 经营者/法人其他类型证件信息
	NeedAccountInfo      bool                                `json:"need_account_info"`      // 是否填写结算银行账户
	AccountInfo          *EcommerceApplyAccountInfo          `json:"account_info"`           // 结算银行账户
	ContactInfo          *EcommerceApplyContactInfo          `json:"contact_info"`           // 超级管理员
	SalesSceneInfo       *EcommerceApplySalesSceneInfo       `json:"sales_scene_info"`       // 店铺信息
	MerchantShortname    string                              `json:"merchant_shortname"`     // 商户简称
	Qualifications       string                              `json:"qualifications"`         // 特殊资质
	BusinessAdditionPics string                              `json:"business_addition_pics"` // 补充材料
	BusinessAdditionDesc string                              `json:"business_addition_desc"` // 补充说明
}

// EcommerceApplyBusinessLicenseInfo 二级商户进件-营业执照/登记证书信息
type EcommerceApplyBusinessLicenseInfo struct {
	BusinessLicenseCopy   string `json:"business_license_copy"`   // 证件扫描件
	BusinessLicenseNumber string `json:"business_license_number"` // 证件注册号
	MerchantName          string `json:"merchant_name"`           // 商户名称
	LegalPerson           string `json:"legal_person"`            // 经营者/法定代表人姓名
	CompanyAddress        string `json:"company_address	"`        // 注册地址
	BusinessTime          string `json:"business_time"`           // 营业期限
}

// EcommerceApplyOrganizationCertInfo 二级商户进件-组织机构代码证信息
type EcommerceApplyOrganizationCertInfo struct {
	OrganizationCopy   string `json:"organization_copy"`   // 组织机构代码证照片
	OrganizationNumber string `json:"organization_number"` // 组织机构代码
	OrganizationTime   string `json:"organization_time"`   // 组织机构代码有效期限
}

// EcommerceApplyIDCardInfo 二级商户进件-经营者/法人身份证信息
type EcommerceApplyIDCardInfo struct {
	IDCardCopy      string `json:"id_card_copy"`       // 身份证人像面照片
	IDCardNational  string `json:"id_card_national"`   // 身份证国徽面照片
	IDCardName      string `json:"id_card_name"`       // 身份证姓名
	IDCardNumber    string `json:"id_card_number"`     // 身份证号码
	IDCardValidTime string `json:"id_card_valid_time"` // 身份证有效期限
}

// EcommerceApplyIDDocInfo 二级商户进件-经营者/法人其他类型证件信息
type EcommerceApplyIDDocInfo struct {
	IDDocName    string `json:"id_doc_name"`    // 证件姓名
	IDDocNumber  string `json:"id_doc_number"`  // 证件号码
	IDDocCopy    string `json:"id_doc_copy"`    // 证件照片
	DocPeriodEnd string `json:"doc_period_end"` // 证件结束日期
}

// EcommerceApplyAccountInfo 二级商户进件-结算银行信息
type EcommerceApplyAccountInfo struct {
	BankAccountType string `json:"bank_account_type"` // 账户类型
	AccountBank     string `json:"account_bank"`      // 开户银行
	AccountName     string `json:"account_name"`      // 开户名称
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
	BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
	BankName        string `json:"bank_name"`         // 开户银行全称
	AccountNumber   string `json:"account_number"`    // 银行账号
}

// EcommerceApplyContactInfo 二级商户进件-超级管理员信息
type EcommerceApplyContactInfo struct {
	ContactType         string `json:"contact_type"`           // 超级管理员类型
	ContactName         string `json:"contact_name"`           // 超级管理员姓名
	ContactIDCardNumber string `json:"contact_id_card_number"` // 超级管理员身份证件号码
	MobilePhone         string `json:"mobile_phone"`           // 超级管理员手机
	ContactEmail        string `json:"contact_email"`          // 超级管理员邮箱
}

// EcommerceApplySalesSceneInfo 二级商户进件-店铺信息
type EcommerceApplySalesSceneInfo struct {
	StoreName           string `json:"store_name"`             // 店铺名称
	StoreURL            string `json:"store_url"`              // 店铺链接
	StoreQrCode         string `json:"store_qr_code"`          // 店铺二维码
	MiniProgramSubAPPID string `json:"mini_program_sub_appid"` // 小程序APPID
}

// EcommerceApplyResp 二级商户进件返回参数
type EcommerceApplyResp struct {
	ApplymentID  int64  `json:"applyment_id"`   // 微信支付申请单号
	OutRequestNo string `json:"out_request_no"` // 业务申请编号
}

// EcommerceApplyQuery 二级商户进件查询
type EcommerceApplyQuery struct {
	ApplymentID string `json:"applyment_id"` // 微信支付申请单号
}

// EcommerceApplyQueryResp 二级商户进件查询
type EcommerceApplyQueryResp struct {
	ApplymentState     string `json:"applyment_state"`      // 申请状态
	ApplymentStateDesc string `json:"applyment_state_desc"` // 申请状态描述
	SignURL            string `json:"sign_url"`             // 签约链接
	SubMchid           string `json:"sub_mchid"`            // 电商平台二级商户号
	AccountValidation  struct {
		AccountName              int64  `json:"account_name"`               // 付款户名
		AccountNo                string `json:"account_no"`                 // 付款卡号
		PayAmount                string `json:"pay_amount"`                 // 汇款金额
		DestinationAccountNumber string `json:"destination_account_number"` // 收款卡号
		DestinationAccountName   string `json:"destination_account_name"`   // 收款户名
		DestinationAccountBank   string `json:"destination_account_bank"`   // 开户银行
		City                     string `json:"city"`                       // 省市信息
		Remark                   string `json:"remark"`                     // 备注信息
		Deadline                 string `json:"deadline"`                   // 汇款截止日期
	} `json:"account_validation"` // 汇款账户验证信息
	AuditDetail []struct {
		ParamName    string `json:"param_name"`    // 参数名称
		RejectReason string `json:"reject_reason"` // 驳回原因
	} `json:"audit_detail"` // 驳回原因详情
	LegalValidationURL string `json:"legal_validation_url"` // 法人验证链接
	OutRequestNo       string `json:"out_request_no"`       // 业务申请编号
	ApplymentID        int64  `json:"applyment_id"`         // 微信支付申请单号
}

// ProfitSharingReceiversAdd 添加分账接收方
type ProfitSharingReceiversAdd struct {
	Appid         string `json:"appid"`          // 公众账号ID
	Type          string `json:"type"`           // 接收方类型
	Account       string `json:"account"`        // 接收方账号
	Name          string `json:"name"`           // 接收方名称
	EncryptedName string `json:"encrypted_name"` // 接收方名称的密文
	RelationType  string `json:"relation_type"`  // 与分账方的关系类型
}

// Apply 请求分账参数
type ProfitSharingApply struct {
	Appid         string                        `json:"appid"`          // 公众账号ID
	SubMchid      string                        `json:"sub_mchid"`      // 二级商户号
	TransactionID string                        `json:"transaction_id"` // 微信订单号
	OutOrderNo    string                        `json:"out_order_no"`   // 商户分账单号
	Receivers     []*ProfitSharingApplyReceiver `json:"receivers"`      // 分账接收方列表
	Finish        bool                          `json:"finish"`         // 是否分账完成
}

// ProfitSharingApplyReceiver 请求分账-分账接收方列表
type ProfitSharingApplyReceiver struct {
	Type            string `json:"type"`             // 分账接收方类型
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号
	Amount          int    `json:"amount"`           // 分账金额
	Description     string `json:"description"`      // 分账描述
	ReceiverName    string `json:"receiver_name"`    // 分账姓名
}

// Query 分账-查询分账结果
type ProfitSharingQuery struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
}

// FinishOrder 分账-完结分账
type ProfitSharingFinishOrder struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	Description   string `json:"description"`    // 分账描述
}

// RefundApply 申请退款
type RefundApply struct {
	SubMchid string `json:"sub_mchid"` // 二级商户号
	SpAppid  string `json:"sp_appid"`  // 电商平台APPID
	//SubAppid      string             `json:"sub_appid"`      // 二级商户APPID
	TransactionID string             `json:"transaction_id"` // 微信订单号
	OutOrderNo    string             `json:"out_order_no"`   // 商户订单号
	OutRefundNo   string             `json:"out_refund_no"`  // 商户退款单号
	Reason        string             `json:"reason"`         // 退款原因
	Amount        *RefundApplyAmount `json:"amount"`         // 订单金额
	NotifyURL     string             `json:"notify_url"`     // 退款结果回调URL
}

// RefundApplyAmount 申请退款-订单金额
type RefundApplyAmount struct {
	Refund   int    `json:"refund"`   // 退款金额
	Total    int    `json:"total"`    // 原订单金额
	Currency string `json:"currency"` // 退款币种
}

// RefundQuery 退款查询
type RefundQuery struct {
	RefundID string `json:"refund_id"` // 微信退款单号
	SubMchid string `json:"sub_mchid"` // 二级商户号
}

// RefundQueryByRefundNo 退款查询
type RefundQueryByRefundNo struct {
	OutRefundNo string `json:"out_refund_no"` // 退款单号
	SubMchid string `json:"sub_mchid"` // 二级商户号
}

// BalanceSubMch 二级商户余额查询
type BalanceSubMch struct {
	SubMchid string `json:"sub_mchid"` // 二级商户号
}

// BalanceSubMchResp 二级商户余额查询
type BalanceSubMchResp struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}

// WithdrawSubMch 二级商户余额提现
type WithdrawSubMch struct {
	SubMchid     string `json:"sub_mchid"`      // 二级商户号
	OutRequestNo string `json:"out_request_no"` // 商户提现单号
	Amount       int    `json:"amount"`         // 提现金额
	Remark       string `json:"remark"`         // 提现备注
	BankMemo     string `json:"bank_memo"`      // 银行附言
}

// WithdrawSubMchQuery 二级商户提现单号查询
type WithdrawSubMchQuery struct {
	SubMchid   string `json:"sub_mchid"`   // 二级商户号
	WithdrawID string `json:"withdraw_id"` // 微信支付提现单号
}

// ProfitSharingReceiversAddResp 订单查询返回参数
type ProfitSharingReceiversAddResp struct {
	Type    string `json:"type"`    // 接收方类型
	Account string `json:"account"` // 接收方账号
}

// ProfitSharingApplyResp .
type ProfitSharingApplyResp struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderID       string `json:"order_id"`       // 微信分账单号
}

// ProfitSharingQueryResp 分账查询
type ProfitSharingQueryResp struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderID       string `json:"order_id"`       // 微信分账单号
	Status        string `json:"status"`         // 分账状态
	Receivers     []struct {
		ReceiverMchid   string `json:"receiver_mchid"`   // 分账接收商户号
		Amount          int    `json:"amount"`           // 分账金额
		Description     string `json:"description"`      // 分账描述
		Result          string `json:"result"`           // 分账结果
		FinishTime      string `json:"finish_time"`      // 完成时间
		FailReason      string `json:"fail_reason"`      // 分账失败原因
		Type            string `json:"type"`             // 分账接收方类型
		ReceiverAccount string `json:"receiver_account"` // 分账接收方账号
	} `json:"receivers"` // 分账接收方列表
	CloseReason       string `json:"close_reason"`       // 关单原因
	FinishAmount      int    `json:"finish_amount"`      // 分账完结金额
	FinishDescription string `json:"finish_description"` // 分账完结描述
}

// ProfitSharingFinishOrderResp 完结分账返回结果
type ProfitSharingFinishOrderResp struct {
	SubMchid      string `json:"sub_mchid"`      // 二级商户号
	TransactionID string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	OrderID       string `json:"order_id"`       // 微信分账单号
}

// RefundApplyResp 退款申请
type RefundApplyResp struct {
	RefundID    string `json:"refund_id"`     // 微信退款单号
	OutRefundNo string `json:"out_refund_no"` // 商户退款单号
	CreateTime  string `json:"create_time"`   // 退款创建时间
	Amount      struct {
		Refund         int    `json:"refund"`          // 退款金额
		PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
		DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
		Currency       string `json:"currency"`        // 退款币种
	} `json:"amount"` // 订单金额
	PromotionDetail []struct {
		PromotionID  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
	} `json:"promotion_detail"` // 优惠退款详情
}

// RefundQueryResp 退款申请查询
type RefundQueryResp struct {
	RefundID            string `json:"refund_id"`             // 微信退款单号
	OutRefundNo         string `json:"out_refund_no"`         // 商户退款单号
	TransactionID       string `json:"transaction_id"`        // 微信订单号
	OutTradeNo          string `json:"out_trade_no"`          // 商户订单号
	Channel             string `json:"channel"`               // 退款渠道
	UserReceivedAccount string `json:"user_received_account"` // 退款入账账号
	SuccessTime         string `json:"success_time"`          // 退款成功时间
	CreateTime          string `json:"create_time"`           // 退款创建时间
	Status              string `json:"status"`                // 退款状态
	Amount              struct {
		Refund         int    `json:"refund"`          // 退款金额
		PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
		DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
		Currency       string `json:"currency"`        // 退款币种
	} `json:"amount"` // 订单金额
	PromotionDetail []struct {
		PromotionID  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
	} `json:"promotion_detail"` // 优惠退款详情
}

// WithdrawSubMchResp 二级商户提现
type WithdrawSubMchResp struct {
	SubMchid     string `json:"sub_mchid"`      // 二级商户号
	WithdrawID   string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo string `json:"out_request_no"` // 商户提现单号s
}

// WithdrawSubMchQueryResp 二级商户提现查询
type WithdrawSubMchQueryResp struct {
	SubMchid     string `json:"sub_mchid"`      // 二级商户号
	SpMchid      string `json:"sp_mchid"`       // 电商平台商户号
	Status       string `json:"status"`         // 提现单状态
	WithdrawID   string `json:"withdraw_id"`    // 微信支付提现单号
	OutRequestNo string `json:"out_request_no"` // 商户提现单号
	Amount       int    `json:"amount"`         // 提现金额
	CreateTime   string `json:"create_time"`    // 发起提现时间
	UpdateTime   string `json:"update_time"`    // 提现状态更新时间
	Reason       string `json:"reason"`         // 失败原因
	Remark       string `json:"remark"`         // 提现备注
	BankMemo     string `json:"bank_memo"`      // 银行附言
}

// CommonImageUpload 通用图片上传
type CommonImageUpload struct {
	FilePath string `json:"file_path"`
}

// CommonImageUploadResp 通用图片上传返回
type CommonImageUploadResp struct {
	MediaID string `json:"media_id"`
}

// EcommerceModifySettlement 修改结算信息
type EcommerceModifySettlement struct {
	SubMchid        string `json:"sub_mchid"`         // 特约商户号
	AccountType     string `json:"account_type"`      // 账户类型
	AccountBank     string `json:"account_bank"`      // 开户银行
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
	BankName        string `json:"bank_name"`         // 开户银行全称（含支行）
	BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
	AccountNumber   string `json:"account_number"`    // 银行账号
}

// EcommerceModifySettlementBody 修改结算信息
type EcommerceModifySettlementBody struct {
	AccountType     string `json:"account_type"`      // 账户类型
	AccountBank     string `json:"account_bank"`      // 开户银行
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
	//BankName        string `json:"bank_name"`         // 开户银行全称（含支行）
	//BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
	AccountNumber   string `json:"account_number"`    // 银行账号
}

// EcommerceQuerySettlement 查询结算信息
type EcommerceQuerySettlement struct {
	SubMchid string `json:"sub_mchid"` // 特约商户号
}

// EcommerceQuerySettlementResp .
type EcommerceQuerySettlementResp struct {
	AccountType   string `json:"account_type"`   // 账户类型
	AccountBank   string `json:"account_bank"`   // 开户银行
	BankName      string `json:"bank_name"`      // 开户银行全称（含支行）
	BankBranchID  string `json:"bank_branch_id"` // 开户银行联行号
	AccountNumber string `json:"account_number"` // 银行账号
	VerifyResult  string `json:"verify_result"`  // 汇款验证结果
}

// ProfitSharingReturnOrders 分账回退
type ProfitSharingReturnOrders struct {
	SubMchid string `json:"sub_mchid"` // 二级商户号
	//OrderID     string `json:"order_id"`      // 微信分账单号
	OutOrderNo  string `json:"out_order_no"`  // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
	ReturnMchid string `json:"return_mchid"`  // 回退商户号
	Amount      int    `json:"amount"`        // 回退金额
	Description string `json:"description"`   // 回退描述
}
// ProfitSharingReturnOrdersResp .
type ProfitSharingReturnOrdersResp struct {
	SubMchid    string `json:"sub_mchid"`     // 二级商户号
	OrderID     string `json:"order_id"`      // 微信分账单号
	OutOrderNo  string `json:"out_order_no"`  // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
	ReturnMchid string `json:"return_mchid"`  // 回退商户号
	Amount      int    `json:"amount"`        // 回退金额
	ReturnNo    string `json:"return_no"`     // 微信回退单号
	Result      string `json:"result"`        // 回退结果
	FailReason  string `json:"fail_reason"`   // 失败原因
	FinishTime  string `json:"finish_time"`   // 完成时间
}

// ProfitSharingReturnOrdersQuery 分账回退查询
type ProfitSharingReturnOrdersQuery struct {
	SubMchid string `json:"sub_mchid"` // 二级商户号
	OutOrderNo  string `json:"out_order_no"`  // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
}

//  ProfitSharingReturnOrdersQueryResp .
type ProfitSharingReturnOrdersQueryResp struct {
	SubMchid    string `json:"sub_mchid"`     // 二级商户号
	OrderID     string `json:"order_id"`      // 微信分账单号
	OutOrderNo  string `json:"out_order_no"`  // 商户分账单号
	OutReturnNo string `json:"out_return_no"` // 商户回退单号
	ReturnMchid string `json:"return_mchid"`  // 回退商户号
	Amount      int    `json:"amount"`        // 回退金额
	ReturnNo    string `json:"return_no"`     // 微信回退单号
	Result      string `json:"result"`        // 回退结果
	FailReason  string `json:"fail_reason"`   // 失败原因
	FinishTime  string `json:"finish_time"`   // 完成时间
}

// ProfitSharingLeftOrderAmount 查询订单剩余待分金额
type ProfitSharingLeftOrderAmount struct {
	TransactionID string `json:"transaction_id"` // 订单号
}

//  ProfitSharingLeftOrderAmountResp .
type ProfitSharingLeftOrderAmountResp struct {
	TransactionID string `json:"transaction_id"` // 订单号
	UnsplitAmount     string `json:"unsplit_amount"`      // 订单剩余待分金额
}


// Error 错误信息
type Error struct {
	Code    string `json:"code"`    // 详细错误码
	Message string `json:"message"` // 错误描述
	Detail  struct {
		Field    string `json:"field"` // 错误参数的位置
		Value    string `json:"value"` // 错误的值
		Issue    string `json:"issue"` // 具体错误原因
		Location string `json:"location"`
	} `json:"detail"` // 具体错误信息
}

// CertCertificates 平台证书列表
type CertCertificatesResp struct {
	List []CertCertificatesListResp `json:"list"`
}

// CertCertificatesListResp 平台证书-单证书信息
type CertCertificatesListResp struct {
	SerialNo  string `json:"serial_no"`  // 证书序列号
	PublicKey string `json:"public_key"` // 微信证书公钥
}

// CipherResp 微信返回的加密结构体
type CipherResp struct {
	Data []struct {
		SerialNo           string `json:"serial_no"`
		EffectiveTime      string `json:"effective_time"`
		ExpireTime         string `json:"expire_time"`
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			Nonce          string `json:"nonce"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
	} `json:"data"`
}

// HTTP CODE不等于200，或204时的错误返回参数
type ErrorResponse struct {
	Code string
	Message string
	Detail struct {
		Field string
		Value string
		Issue string
		Location string
	}
}
