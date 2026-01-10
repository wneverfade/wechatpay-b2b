package model

// RefundStatus 表示退款状态。
type RefundStatus string

const (
	// RefundInit 初始化/受理中
	RefundInit RefundStatus = "REFUND_INIT"
	// RefundProcessing 处理中
	RefundProcessing RefundStatus = "REFUND_PROCESSING"
	// RefundSuccess 成功
	RefundSuccess RefundStatus = "REFUND_SUCCESS"
	// RefundFail 失败
	RefundFail RefundStatus = "REFUND_FAIL"
)
