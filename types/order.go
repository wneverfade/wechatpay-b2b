package types

import "github.com/enjoy322/wechatpay-b2b/model"

// CloseOrderRequest 关闭订单请求参数。
type CloseOrderRequest struct {
	Mchid      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

// CloseOrderResponse 关闭订单返回参数。
type CloseOrderResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetOrderRequest 查询订单请求参数。
type GetOrderRequest struct {
	Mchid      string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

// GetOrderResponse 查询订单返回参数。
type GetOrderResponse struct {
	AppID        string             `json:"appid"`
	Mchid        string             `json:"mchid"`
	OutTradeNo   string             `json:"out_trade_no"`
	OrderID      string             `json:"order_id"`
	PayStatus    model.PayStatus    `json:"pay_status"`
	PayTime      string             `json:"pay_time,omitempty"`
	Attach       string             `json:"attach,omitempty"`
	RefundStatus model.RefundStatus `json:"refund_status,omitempty"`
	ErrCode      int                `json:"errcode,omitempty"`
	ErrMsg       string             `json:"errmsg,omitempty"`
}
