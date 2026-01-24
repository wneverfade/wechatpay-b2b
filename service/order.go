package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/enjoy322/wechatpay-b2b/client"
	"github.com/enjoy322/wechatpay-b2b/types"
)

// OrderService 处理订单、退款与支付参数构建相关调用。
type OrderService interface {
	// CloseOrder 关闭订单。
	CloseOrder(ctx context.Context, req types.CloseOrderRequest, appKey string) (*types.CloseOrderResponse, error)
	// GetOrder 查询订单。
	GetOrder(ctx context.Context, req types.GetOrderRequest, appKey string) (*types.GetOrderResponse, error)
	// CreateRefund 发起退款。
	CreateRefund(ctx context.Context, req types.RefundRequest, appKey string) (*types.RefundResponse, error)
	// GetRefund 查询退款。
	GetRefund(ctx context.Context, req types.GetRefundRequest, appKey string) (*types.GetRefundResponse, error)
	// BuildPaymentParams 生成单订单支付参数，用于小程序 wx.requestCommonPayment。
	// BuildPaymentParams(ctx context.Context, req types.Order, sessionKey string, appKey string) (*types.CommonPaymentParams, error)
	// BuildCombinedPaymentParams 生成合单支付参数，用于小程序 wx.requestCommonPayment。
	BuildCombinedPaymentParams(ctx context.Context, req types.CombinedPaymentSignData, sessionKey string) (*types.CommonPaymentParams, error)
}

type orderService struct {
	client *client.Client
}

const (
	closeOrderURI           = "/retail/B2b/closeb2border"
	getOrderURI             = "/retail/B2b/getorder"
	createRefundURI         = "/retail/B2b/refund"
	getRefundURI            = "/retail/B2b/getrefund"
	requestCommonPaymentURI = "requestCommonPayment"
	paymentModeGoods        = "retail_pay_goods"
	paymentModeCombined     = "retail_pay_combined_goods"
)

// NewOrderService 创建订单服务。
func NewOrderService(c *client.Client) OrderService {
	return &orderService{client: c}
}

// CloseOrder 关闭订单。
func (s *orderService) CloseOrder(ctx context.Context, req types.CloseOrderRequest, appKey string) (*types.CloseOrderResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(closeOrderURI, body, appKey)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.CloseOrderResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetOrder 查询订单。
func (s *orderService) GetOrder(ctx context.Context, req types.GetOrderRequest, appKey string) (*types.GetOrderResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(getOrderURI, body, appKey)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.GetOrderResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// CreateRefund 发起退款。
func (s *orderService) CreateRefund(ctx context.Context, req types.RefundRequest, appKey string) (*types.RefundResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.OutRefundNo == "" {
		return nil, errors.New("out_refund_no is required")
	}
	if req.RefundAmount <= 0 {
		return nil, errors.New("refund_amount is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(createRefundURI, body, appKey)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.RefundResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetRefund 查询退款。
func (s *orderService) GetRefund(ctx context.Context, req types.GetRefundRequest, appKey string) (*types.GetRefundResponse, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.OutRefundNo == "" && req.RefundID == "" {
		return nil, errors.New("out_refund_no or refund_id is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuthAndSig(getRefundURI, body, appKey)

	resp, err := s.client.Do(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("wechat api http status %d: %s", resp.StatusCode, string(raw))
	}

	var out types.GetRefundResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// BuildPaymentParams 生成单订单支付参数，用于小程序 wx.requestCommonPayment。
func (s *orderService) BuildPaymentParams(ctx context.Context, req types.Order, sessionKey string, appKey string) (*types.CommonPaymentParams, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}
	if sessionKey == "" {
		return nil, errors.New("sessionKey is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return &types.CommonPaymentParams{
		SignData:  string(body),
		Mode:      paymentModeGoods,
		PaySig:    s.client.GetPaySig(requestCommonPaymentURI, body, appKey),
		Signature: s.client.GetUserSignature(body, sessionKey),
	}, nil
}

// BuildCombinedPaymentParams 生成合单支付参数，用于小程序 wx.requestCommonPayment。
func (s *orderService) BuildCombinedPaymentParams(ctx context.Context, req types.CombinedPaymentSignData, sessionKey string) (*types.CommonPaymentParams, error) {
	if s.client == nil {
		return nil, errors.New("client is nil")
	}

	if sessionKey == "" {
		return nil, errors.New("sessionKey is empty")
	}
	if len(req.CombinedOrderList) == 0 {
		return nil, errors.New("combined_order_list is required")
	}

	type paySigItem struct {
		Mchid  string `json:"mchid"`
		PaySig string `json:"paysig"`
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	paySigItems := make([]paySigItem, 0, len(req.CombinedOrderList))
	for _, order := range req.CombinedOrderList {
		if order.Mchid == "" {
			return nil, errors.New("mchid is required")
		}
		paySigPer := s.client.GetPaySig(requestCommonPaymentURI, body, order.AppKey)

		paySigItems = append(paySigItems, paySigItem{
			Mchid:  order.Mchid,
			PaySig: paySigPer,
		})
	}

	paySigBytes, err := json.Marshal(paySigItems)
	if err != nil {
		return nil, err
	}

	return &types.CommonPaymentParams{
		SignData:  string(body),
		Mode:      paymentModeCombined,
		PaySig:    string(paySigBytes),
		Signature: s.client.GetUserSignature(body, sessionKey),
	}, nil
}
