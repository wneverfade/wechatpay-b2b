package types

import "github.com/enjoy322/wechatpay-b2b/model"

// ProfitSharingRequest 请求分账参数。
type ProfitSharingRequest struct {
	Mchid          string                  `json:"mchid"`                     // 微信商户号
	TransactionID  string                  `json:"transaction_id"`            // 微信支付订单号
	OutOrderNo     string                  `json:"out_order_no"`              // 商户分账订单号
	UnfreezeUnpaid bool                    `json:"unfreeze_unpaid,omitempty"` // 是否解冻未分账金额
	Receivers      []model.ProfitReceiver  `json:"receivers"`                 // 分账接收方列表
	Description    string                  `json:"description,omitempty"`     // 分账描述
}

// ProfitSharingResponse 请求分账返回参数。
type ProfitSharingResponse struct {
	OrderID     string `json:"order_id"`               // 分账订单号
	OutOrderNo  string `json:"out_order_no"`           // 商户分账订单号
	TransactionID string `json:"transaction_id"`       // 微信支付订单号
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// QueryProfitSharingRequest 查询分账订单参数。
type QueryProfitSharingRequest struct {
	Mchid       string `json:"mchid"`                 // 微信商户号
	TransactionID string `json:"transaction_id"`      // 微信支付订单号
	OutOrderNo  string `json:"out_order_no,omitempty"`// 商户分账订单号
}

// QueryProfitSharingResponse 查询分账订单返回参数。
type QueryProfitSharingResponse struct {
	OrderID          string                  `json:"order_id"`                    // 分账订单号
	OutOrderNo       string                  `json:"out_order_no"`                // 商户分账订单号
	TransactionID    string                  `json:"transaction_id"`              // 微信支付订单号
	Status           model.ProfitStatus      `json:"status"`                      // 分账状态
	Receivers        []model.ProfitReceiver  `json:"receivers"`                   // 分账接收方列表
	Amount           int64                   `json:"amount"`                      // 分账金额
	ErrCode          int                     `json:"errcode"`
	ErrMsg           string                  `json:"errmsg"`
}

// ProfitSharingFinishRequest 分账完结参数。
type ProfitSharingFinishRequest struct {
	Mchid          string `json:"mchid"`                     // 微信商户号
	TransactionID  string `json:"transaction_id"`            // 微信支付订单号
	OutOrderNo     string `json:"out_order_no"`              // 商户分账订单号
	Description    string `json:"description,omitempty"`     // 完结描述
}

// ProfitSharingFinishResponse 分账完结返回参数。
type ProfitSharingFinishResponse struct {
	OrderID       string `json:"order_id"`          // 分账订单号
	OutOrderNo    string `json:"out_order_no"`      // 商户分账订单号
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

// ProfitSharingReturnRequest 分账回退参数。
type ProfitSharingReturnRequest struct {
	Mchid          string `json:"mchid"`                     // 微信商户号
	TransactionID  string `json:"transaction_id"`            // 微信支付订单号
	OutOrderNo     string `json:"out_order_no"`              // 商户分账订单号
	OutReturnNo    string `json:"out_return_no"`             // 商户回退订单号
	ReturnAmount   int64  `json:"return_amount"`             // 回退金额
	Receiver       model.ProfitReceiver `json:"receiver"`           // 分账接收方
	Description    string `json:"description,omitempty"`     // 描述
}

// ProfitSharingReturnResponse 分账回退返回参数。
type ProfitSharingReturnResponse struct {
	ReturnID      string `json:"return_id"`       // 回退订单号
	OutReturnNo   string `json:"out_return_no"`   // 商户回退订单号
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

// QueryProfitSharingReturnRequest 查询分账回退参数。
type QueryProfitSharingReturnRequest struct {
	Mchid       string `json:"mchid"`                 // 微信商户号
	OutReturnNo string `json:"out_return_no"`         // 商户回退订单号
	ReturnID    string `json:"return_id,omitempty"`   // 回退订单号
}

// QueryProfitSharingReturnResponse 查询分账回退返回参数。
type QueryProfitSharingReturnResponse struct {
	ReturnID       string               `json:"return_id"`        // 回退订单号
	OutReturnNo    string               `json:"out_return_no"`    // 商户回退订单号
	OrderID        string               `json:"order_id"`         // 分账订单号
	Amount         int64                `json:"amount"`           // 回退金额
	Status         model.ProfitReturnStatus `json:"status"`     // 回退状态
	ErrCode        int                  `json:"errcode"`
	ErrMsg         string               `json:"errmsg"`
}
