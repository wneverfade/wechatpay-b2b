package model

// BalanceType 表示账户资金类型。
type BalanceType string

const (
	// BalanceAvailable 可提现金额
	BalanceAvailable BalanceType = "BALANCE_TYPE_AVAILABLE"
	// BalanceFrozen 待结算金额
	BalanceFrozen BalanceType = "BALANCE_TYPE_FROZEN"
)

// BalanceInfo 表示账户余额条目。
type BalanceInfo struct {
	BalanceType BalanceType `json:"balance_type"`
	Amount      string      `json:"amount"`
	Currency    string      `json:"currency"`
}

// WithdrawStatus 表示提现状态。
type WithdrawStatus string

const (
	// WithdrawInit 初始化
	WithdrawInit WithdrawStatus = "WITHDRAW_INIT"
	// WithdrawProcess 进行中
	WithdrawProcess WithdrawStatus = "WITHDRAW_PROCESSING"
	// WithdrawSuccess 成功
	WithdrawSuccess WithdrawStatus = "WITHDRAW_SUCC"
	// WithdrawFail 失败
	WithdrawFail WithdrawStatus = "WITHDRAW_FAIL"
	// WithdrawRefund 退票/回退
	WithdrawRefund WithdrawStatus = "WITHDRAW_REFUND"
)
