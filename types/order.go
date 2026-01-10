package types

import "github.com/enjoy322/wechatpay-b2b/model"

// CloseOrderRequest 关闭订单请求参数。
type CloseOrderRequest struct {
	Mchid      string `json:"mchid"`                  // 微信商户号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号
	OrderID    string `json:"order_id,omitempty"`     // B2b支付订单号
}

// CloseOrderResponse 关闭订单返回参数。
type CloseOrderResponse struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

// GetOrderRequest 查询订单请求参数。
type GetOrderRequest struct {
	Mchid      string `json:"mchid"`                  // 微信商户号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号
	OrderID    string `json:"order_id,omitempty"`     // B2b支付订单号
}

// GetOrderResponse 查询订单返回参数。
type GetOrderResponse struct {
	AppID                 string             `json:"appid"`                             // 小程序ID
	Mchid                 string             `json:"mchid"`                             // 微信商户号
	OutTradeNo            string             `json:"out_trade_no"`                      // 商户订单号
	OrderID               string             `json:"order_id"`                          // B2b支付订单号
	PayStatus             model.PayStatus    `json:"pay_status"`                        // 订单状态
	PayTime               string             `json:"pay_time,omitempty"`                // 支付完成时间
	Attach                string             `json:"attach,omitempty"`                  // 附加数据
	PayerOpenID           string             `json:"payer_openid"`                      // 支付者
	Amount                model.Amount       `json:"amount"`                            // 订单金额
	WxPayTransactionID    string             `json:"wxpay_transaction_id,omitempty"`    // 微信支付订单号
	Env                   int                `json:"env"`                               // 订单环境
	SettleStatus          uint32             `json:"settle_status"`                     // 结算状态
	SettleFinishTime      string             `json:"settle_finish_time,omitempty"`      // 结算完成时间
	PlatformProfitPercent uint64             `json:"platform_profit_percent,omitempty"` // 技术服务费率
	PlatformProfitFee     uint64             `json:"platform_profit_fee,omitempty"`     // 技术服务费
	BankType              string             `json:"bank_type,omitempty"`               // 银行类型
	RefundStatus          model.RefundStatus `json:"refund_status,omitempty"`           // 退款状态
	ErrCode               int                `json:"errcode,omitempty"`                 // 错误码
	ErrMsg                string             `json:"errmsg,omitempty"`                  // 错误信息
}
