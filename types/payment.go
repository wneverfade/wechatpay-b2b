package types

// Order 表示单笔微信支付下单的参数。
type Order struct {
	Mchid        string       `json:"mchid"`
	OutTradeNo   string       `json:"out_trade_no"`
	Description  string       `json:"description"`
	Amount       Amount       `json:"amount"`
	Attach       string       `json:"attach"`        // 附加数据-选填
	ProductInfo  *ProductInfo `json:"product_info"`  // 商品信息-选填
	DeliveryType uint32       `json:"delivery_type"` // 配送类型-选填
	Env          uint32       `json:"env"`           // 环境
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

// Amount 金额明细。
type Amount struct {
	ProductAmount int64  `json:"product_amount,omitempty"`
	Freight       int64  `json:"freight,omitempty"`
	Discount      int64  `json:"discount,omitempty"`
	OtherFee      int64  `json:"other_fee,omitempty"`
	OrderAmount   int64  `json:"order_amount"`
	PayerAmount   int64  `json:"payer_amount,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// CombinedPaymentSignData 合单支付的 signData 结构。
type CombinedPaymentSignData struct {
	Env               uint32           `json:"env"`
	CombinedOrderList []*CombinedOrder `json:"combined_order_list"`
}

// CommonPaymentParams 返回给小程序 wx.requestCommonPayment 的参数集。
type CommonPaymentParams struct {
	SignData  string `json:"signData"`
	Mode      string `json:"mode"`
	PaySig    string `json:"paySig"`
	Signature string `json:"signature"`
}
