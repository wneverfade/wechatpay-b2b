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

// MerchantService 处理商户信息与余额相关接口调用。
type MerchantService interface {
	// GetMerchantInfo 获取小程序下所有商户的信息。
	GetMerchantInfo(ctx context.Context, req types.GetMerchantInfoRequest) (*types.GetMerchantInfoResponse, error)
	// GetMerchantAppKey 查询商户的 appKey。
	GetMerchantAppKey(ctx context.Context, req types.GetMerchantAppKeyRequest) (*types.GetMerchantAppKeyResponse, error)
	// GetBalance 查询账户余额。
	GetBalance(ctx context.Context, req types.BalanceRequest, appKey string) (*types.BalanceResponse, error)
	// Withdraw 发起提现。
	Withdraw(ctx context.Context, req types.WithdrawRequest, appKey string) (*types.WithdrawResponse, error)
	// QueryWithdraw 查询提现状态。
	QueryWithdraw(ctx context.Context, req types.QueryWithdrawRequest, appKey string) (*types.QueryWithdrawResponse, error)
}

type merchantService struct {
	client *client.Client
}

const (
	getMerchantInfoURI   = "/retail/B2b/getmchinfo"
	getMerchantAppKeyURI = "/retail/B2b/getappkey"
	getMchBalanceURI     = "/retail/B2b/getmchbalance"
	withdrawURI          = "/retail/B2b/withdraw"
	queryWithdrawURI     = "/retail/B2b/querywithdraw"
)

// NewMerchantService 创建商户信息服务。
func NewMerchantService(c *client.Client) MerchantService {
	return &merchantService{client: c}
}

// GetMerchantInfo 获取小程序下所有商户的信息。
func (s *merchantService) GetMerchantInfo(ctx context.Context, req types.GetMerchantInfoRequest) (*types.GetMerchantInfoResponse, error) {
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

	uri := s.client.BuildURIWithAuth(getMerchantInfoURI)

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

	var out types.GetMerchantInfoResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetMerchantAppKey 查询商户的 appKey。
func (s *merchantService) GetMerchantAppKey(ctx context.Context, req types.GetMerchantAppKeyRequest) (*types.GetMerchantAppKeyResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if s.client.GetAccessToken() == "" {
		return nil, errors.New("accessToken is empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	uri := s.client.BuildURIWithAuth(getMerchantAppKeyURI)

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

	var out types.GetMerchantAppKeyResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// GetBalance 查询账户余额。
func (s *merchantService) GetBalance(ctx context.Context, req types.BalanceRequest, appKey string) (*types.BalanceResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
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

	uri := s.client.BuildURIWithAuthAndSig(getMchBalanceURI, body, appKey)

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

	var out types.BalanceResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// Withdraw 发起提现。
func (s *merchantService) Withdraw(ctx context.Context, req types.WithdrawRequest, appKey string) (*types.WithdrawResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.WithdrawAmount <= 0 {
		return nil, errors.New("withdraw_amount is required")
	}
	if req.OutWithdrawNo == "" {
		return nil, errors.New("out_withdraw_no is required")
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

	uri := s.client.BuildURIWithAuthAndSig(withdrawURI, body, appKey)

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

	var out types.WithdrawResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}

// QueryWithdraw 查询提现状态。
func (s *merchantService) QueryWithdraw(ctx context.Context, req types.QueryWithdrawRequest, appKey string) (*types.QueryWithdrawResponse, error) {
	if s == nil || s.client == nil {
		return nil, errors.New("client is nil")
	}
	if req.Mchid == "" {
		return nil, errors.New("mchid is required")
	}
	if req.OutWithdrawNo == "" {
		return nil, errors.New("out_withdraw_no is required")
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

	uri := s.client.BuildURIWithAuthAndSig(queryWithdrawURI, body, appKey)

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

	var out types.QueryWithdrawResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return &out, fmt.Errorf("wechat api error: errcode=%d errmsg=%s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}
