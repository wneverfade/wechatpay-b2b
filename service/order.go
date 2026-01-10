package service

import (
	"context"
	"errors"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// OrderService 处理订单查询、关单相关调用。
type OrderService struct {
	Client *client.Client
}

func (s *OrderService) CloseOrder(ctx context.Context, req types.CloseOrderRequest) (*types.CloseOrderResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}

// GetOrder 查询订单。
func (s *OrderService) GetOrder(ctx context.Context, req types.GetOrderRequest) (*types.GetOrderResponse, error) {
	if s.Client == nil {
		return nil, errors.New("client is nil")
	}
	return nil, errors.New("not implemented")
}
