package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/signer"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// PaymentService 负责构建小程序支付相关参数。
type PaymentService struct {
	Client *client.Client
}

// BuildCommonPaymentParams 生成小程序 wx.requestCommonPayment 所需参数。
// uri 必须与计算 paySig 的服务端 API 路径一致（例如 /retail/B2b/getorder）。
func (s *PaymentService) BuildCommonPaymentParams(ctx context.Context, uri string, signData any, appKey, sessionKey, mode string) (*types.CommonPaymentParams, error) {
	if appKey == "" {
		return nil, errors.New("appKey is required")
	}
	body, err := json.Marshal(signData)
	if err != nil {
		return nil, err
	}
	return &types.CommonPaymentParams{
		SignData:  string(body),
		Mode:      mode,
		PaySig:    signer.PaySig(uri, body, appKey),
		Signature: signer.UserSignature(body, sessionKey),
	}, nil
}
