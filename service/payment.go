package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// PaymentService 负责构建小程序支付相关参数。
type PaymentService struct {
	Client *client.Client
}

const combinePayURI = "/retail/B2b/combinepay"

// BuildCombinedPaymentParams 生成合单支付参数，用于小程序 wx.requestCommonPayment。
func (s *PaymentService) BuildCombinedPaymentParams(ctx context.Context, req types.CombinedPaymentSignData) (*types.CommonPaymentParams, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	if s.Client.AppKeyProvider == "" {
		return nil, errors.New("appKeyProvider is empty")
	}
	if len(req.CombinedOrderList) == 0 {
		return nil, errors.New("combined_order_list is required")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return &types.CommonPaymentParams{
		SignData:  string(body),
		Mode:      "0",
		PaySig:    s.Client.GetPaySig(combinePayURI, body),
		Signature: s.Client.GetUserSignature(body),
	}, nil
}
