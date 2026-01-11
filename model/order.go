package model

// Env 表示环境。
type Env int

const (
	// EnvProd 生产环境
	EnvProd Env = 0
	// EnvSandbox 沙箱环境
	EnvSandbox Env = 1
)

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
