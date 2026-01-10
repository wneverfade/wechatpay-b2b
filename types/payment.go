package types

import "github.com/enjoy322/wechatpay-b2b/model"

// CommonPaymentSignData 微信支付 signData 结构（普通单）。
type CommonPaymentSignData struct {
	Mchid        string             `json:"mchid"`
	OutTradeNo   string             `json:"out_trade_no"`
	Description  string             `json:"description"`
	Amount       model.Amount       `json:"amount"`
	Attach       string             `json:"attach,omitempty"`
	ProductInfo  *model.ProductInfo `json:"product_info,omitempty"`
	DeliveryType uint32             `json:"delivery_type,omitempty"`
	Env          uint32             `json:"env"`
}

// CombinedPaymentSignData 合单支付的 signData 结构。
type CombinedPaymentSignData struct {
	Env               uint32                `json:"env"`
	CombinedOrderList []model.CombinedOrder `json:"combined_order_list"`
}

// CommonPaymentParams 返回给小程序 wx.requestCommonPayment 的参数集。
type CommonPaymentParams struct {
	SignData  string `json:"signData"`
	Mode      string `json:"mode"`
	PaySig    string `json:"paySig"`
	Signature string `json:"signature"`
}
