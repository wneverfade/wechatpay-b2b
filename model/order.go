package model

// Env 表示环境。
type Env int

const (
	// EnvProd 生产环境
	EnvProd Env = 0
	// EnvSandbox 沙箱环境
	EnvSandbox Env = 1
)

// Amount 表示订单金额明细。
type Amount struct {
	ProductAmount int64  `json:"product_amount,omitempty"`
	Freight       int64  `json:"freight,omitempty"`
	Discount      int64  `json:"discount,omitempty"`
	OtherFee      int64  `json:"other_fee,omitempty"`
	OrderAmount   int64  `json:"order_amount"`
	PayerAmount   int64  `json:"payer_amount,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// ProductInfo 表示商品信息。
type ProductInfo struct {
	SPUID     string `json:"spu_id"`
	SKUID     string `json:"sku_id"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	HeadImg   string `json:"head_img"`
	Category  string `json:"category"`
	SKUAttr   string `json:"sku_attr"`
	OrgPrice  int32  `json:"org_price"`
	SalePrice int32  `json:"sale_price"`
	Quantity  int32  `json:"quantity"`
}

// Order 表示单笔微信支付下单的参数。
type Order struct {
	Mchid        string       `json:"mchid"`
	OutTradeNo   string       `json:"out_trade_no"`
	Description  string       `json:"description"`
	Amount       Amount       `json:"amount"`
	Attach       string       `json:"attach,omitempty"`
	ProductInfo  *ProductInfo `json:"product_info,omitempty"`
	DeliveryType uint32       `json:"delivery_type,omitempty"`
	Env          uint32       `json:"env"`
}

// CombinedOrder 表示合单中的子单参数。
type CombinedOrder struct {
	Mchid        string       `json:"mchid"`
	OutTradeNo   string       `json:"out_trade_no"`
	Description  string       `json:"description"`
	Amount       Amount       `json:"amount"`
	Attach       string       `json:"attach,omitempty"`
	DeliveryType uint32       `json:"delivery_type,omitempty"`
	ProductInfo  *ProductInfo `json:"product_info,omitempty"`
}

// PayStatus 表示支付状态。
type PayStatus string

const (
	// PayStatusInit 订单初始化
	PayStatusInit PayStatus = "ORDER_INIT"
	// PayStatusPrePay 待支付
	PayStatusPrePay PayStatus = "ORDER_PRE_PAY"
	// PayStatusSuccess 支付成功
	PayStatusSuccess PayStatus = "ORDER_PAY_SUCC"
	// PayStatusClosed 已关闭
	PayStatusClosed PayStatus = "ORDER_CLOSE"
	// PayStatusRefunding 退款处理中
	PayStatusRefunding PayStatus = "ORDER_REFUND_PROCESSING"
	// PayStatusRefunded 已退款
	PayStatusRefunded PayStatus = "ORDER_REFUND"
)
