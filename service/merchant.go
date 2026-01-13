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

// MerchantService 处理商户号进件与状态查询相关接口调用。
type MerchantService interface {
	// RegisterMerchant 商户号进件。
	RegisterMerchant(ctx context.Context, req types.RegisterMerchantRequest) (*types.RegisterMerchantResponse, error)
	// GetMerchantOpenStatus 查询商户号开通状态。
	GetMerchantOpenStatus(ctx context.Context, req types.GetMerchantOpenStatusRequest) (*types.GetMerchantOpenStatusResponse, error)
}

type merchantService struct {
	client *client.Client
}

const (
	registerMerchantURI  = "/retail/B2b/retailregistermch"
	getMerchantStatusURI = "/retail/B2b/retailgetmchorder"
)

// NewMerchantService 创建商户号进件服务。
func NewMerchantService(c *client.Client) MerchantService {
	return &merchantService{client: c}
}

// RegisterMerchant 商户号进件。
func (s *merchantService) RegisterMerchant(ctx context.Context, req types.RegisterMerchantRequest) (*types.RegisterMerchantResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(registerMerchantURI)

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

	var out types.RegisterMerchantResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetMerchantOpenStatus 查询商户号开通状态。
func (s *merchantService) GetMerchantOpenStatus(ctx context.Context, req types.GetMerchantOpenStatusRequest) (*types.GetMerchantOpenStatusResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(getMerchantStatusURI)

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

	var out types.GetMerchantOpenStatusResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}
