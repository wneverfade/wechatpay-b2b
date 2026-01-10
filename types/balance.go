package types

import "github.com/enjoy322/wechatpay-b2b/model"

// BalanceRequest 余额查询请求参数。
type BalanceRequest struct {
	Mchid string `json:"mchid"`
}

// BalanceResponse 余额查询返回参数。
type BalanceResponse struct {
	BalanceList []model.BalanceInfo `json:"balance_list"`
	ErrCode     int                 `json:"errcode"`
	ErrMsg      string              `json:"errmsg"`
}

// WithdrawRequest 提现请求参数。
type WithdrawRequest struct {
	Mchid          string `json:"mchid"`
	WithdrawAmount int64  `json:"withdraw_amount"`
	OutWithdrawNo  string `json:"out_withdraw_no"`
}

// WithdrawResponse 提现返回参数。
type WithdrawResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryWithdrawRequest 查询提现状态请求参数。
type QueryWithdrawRequest struct {
	Mchid         string `json:"mchid"`
	OutWithdrawNo string `json:"out_withdraw_no"`
}

// QueryWithdrawResponse 查询提现状态返回参数。
type QueryWithdrawResponse struct {
	OutWithdrawNo  string               `json:"out_withdraw_no"`
	WithdrawAmount int64                `json:"withdraw_amount"`
	Status         model.WithdrawStatus `json:"status"`
	FailReason     string               `json:"fail_reason,omitempty"`
	ErrCode        int                  `json:"errcode"`
	ErrMsg         string               `json:"errmsg"`
}
