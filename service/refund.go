package service

import (
	"context"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// RefundService 处理退款申请与退款查询。
type RefundService struct {
	Client *client.Client
}

// CreateRefund 发起退款。
func (s *RefundService) CreateRefund(ctx context.Context, req types.RefundRequest) (*types.RefundResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// GetRefund 查询退款。
func (s *RefundService) GetRefund(ctx context.Context, req types.GetRefundRequest) (*types.GetRefundResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}
