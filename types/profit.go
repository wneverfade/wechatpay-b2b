package types

import "github.com/enjoy322/wechatpay-b2b/model"

// ProfitSharingRequest 请求分账参数。
type ProfitSharingRequest struct {
	Mchid           string `json:"mchid"`            // 微信商户号
	OutTradeNo      string `json:"out_trade_no"`     // 商户支付订单号
	ProfitFee       int    `json:"profit_fee"`       // 	分账费用	单位:分，不超过支付单本身的金额	是
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型	同添加分账方时填入的内容	是
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号	同添加分账方时填入的内容	是
}

// ProfitSharingResponse 请求分账返回参数。
type ProfitSharingResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryProfitSharingRequest 查询分账订单参数。
type QueryProfitSharingRequest struct {
	Mchid           string `json:"mchid"`            // 微信商户号
	OutTradeNo      string `json:"out_trade_no"`     // 支付单 id	在B2b小程序中下单的订单 id
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型	同添加分账方时填入的内容
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号	同添加分账方时填入的内容
}

// QueryProfitSharingResponse 查询分账订单返回参数。
type QueryProfitSharingResponse struct {
	OrderStatus int    `json:"order_status"` // 分账状态	枚举值： 1：初始化 2：成功 3：失败
	ErrCode     int    `json:"errcode"`      // 错误码
	ErrMsg      string `json:"errmsg"`       // 错误信息
}

// ProfitSharingFinishRequest 分账完结参数。
type ProfitSharingFinishRequest struct {
	Mchid      string `json:"mchid"`        // 微信商户号
	OutOrderNo string `json:"out_order_no"` // 商户分账订单号
}

// ProfitSharingFinishResponse 分账完结返回参数。
type ProfitSharingFinishResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ProfitSharingReturnRequest 分账回退参数。
type ProfitSharingReturnRequest struct {
	Mchid        string `json:"mchid"`         // 微信商户号
	OutTradeNo   string `json:"out_trade_no"`  // 订单 id	在B2b小程序中下单的订单 id
	OutReturnNo  string `json:"out_return_no"` // 退款单 id	在B2b小程序中下单的退款单 id
	PayeeType    string `json:"payee_type"`    // 退款分账方类型	同添加分账方时填入的内容
	PayeeID      string `json:"payee_id"`      // 退款分账方 id	同添加分账方时填入的内容
	RefundAmount int64  `json:"refund_amount"` // 退款金额	退款金额，单位为分
}

// ProfitSharingReturnResponse 分账回退返回参数。
type ProfitSharingReturnResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryProfitSharingReturnRequest 查询分账回退参数。
type QueryProfitSharingReturnRequest struct {
	Mchid       string `json:"mchid"`               // 微信商户号
	OutReturnNo string `json:"out_return_no"`       // 商户回退订单号
	ReturnID    string `json:"return_id,omitempty"` // 回退订单号
}

// QueryProfitSharingReturnResponse 查询分账回退返回参数。
type QueryProfitSharingReturnResponse struct {
	ReturnID    string                   `json:"return_id"`     // 回退订单号
	OutReturnNo string                   `json:"out_return_no"` // 商户回退订单号
	OrderID     string                   `json:"order_id"`      // 分账订单号
	Amount      int64                    `json:"amount"`        // 回退金额
	Status      model.ProfitReturnStatus `json:"status"`        // 回退状态
	ErrCode     int                      `json:"errcode"`
	ErrMsg      string                   `json:"errmsg"`
}

// AddProfitSharingAccountRequest 添加分账方请求参数。
type AddProfitSharingAccountRequest struct {
	ProfitSharingRelationType string `json:"profit_sharing_relation_type"` // 分账接收方关系类型
	PayeeType                 string `json:"payee_type"`                   // 分账接收方类型
	PayeeID                   string `json:"payee_id"`                     // 分账接收方账号
	PayeeName                 string `json:"payee_name"`                   // 分账接收方名称
}

// AddProfitSharingAccountResponse 添加分账方返回参数。
type AddProfitSharingAccountResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DelProfitSharingAccountRequest 删除分账方请求参数。
type DelProfitSharingAccountRequest struct {
	Mchid     string `json:"mchid"`      // 微信商户号
	PayeeType string `json:"payee_type"` // 分账接收方类型
	PayeeID   string `json:"payee_id"`   // 分账接收方账号
}

// DelProfitSharingAccountResponse 删除分账方返回参数。
type DelProfitSharingAccountResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// QueryProfitSharingAccountRequest 查询分账方请求参数。
type QueryProfitSharingAccountRequest struct {
	Offset int `json:"offset"` // 偏移量 0
	Limit  int `json:"limit"`  // 分页大小，默认10
}

// ProfitSharingAccount 分账接收方信息。
type ProfitSharingAccount struct {
	SharingAccountType string `json:"sharing_account_type"` // 分账接收方关系类型
	SharingAccount     string `json:"sharing_account"`      // 分账接收方账号
	AddTime            int64  `json:"add_time"`             // 添加时间
	UpdateTime         int64  `json:"update_time"`          // 更新时间
	Name               string `json:"name"`                 // 分账接收方名称
}

// QueryProfitSharingAccountResponse 查询分账方返回参数。
type QueryProfitSharingAccountResponse struct {
	AccountList []ProfitSharingAccount `json:"account_list"` // 分账接收方列表
	ErrCode     int                    `json:"errcode"`
	ErrMsg      string                 `json:"errmsg"`
}

// QueryProfitSharingRemainAmtRequest 查询分账剩余金额请求参数。
type QueryProfitSharingRemainAmtRequest struct {
	Mchid      string `json:"mchid"`        // 微信商户号
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
}

// QueryProfitSharingRemainAmtResponse 查询分账剩余金额返回参数。
type QueryProfitSharingRemainAmtResponse struct {
	RemainAmount int64  `json:"remain_amount"` // 剩余可分账金额
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}
