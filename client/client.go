package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

// Client 封装访问微信 API 的 HTTP 客户端。
type Client struct {
	BaseURL        string
	TokenProvider  string // access_token
	AppKeyProvider string // appKey（用于计算 pay_sig）
	HTTPClient     *http.Client
}

// Do 向指定 uri（路径，不含完整域名）发起 HTTP 请求。
func (c *Client) Do(ctx context.Context, method, uri string, body []byte) (*http.Response, error) {
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+uri, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return httpClient.Do(req)
}

// GetPaySig 计算 pay_sig，算法为 HMAC-SHA256(appKey, uri+"&"+body)。
func (c *Client) GetPaySig(uri string, body []byte) string {
	return GetPaySig(uri, body, c.AppKeyProvider)
}

// GetUserSignature 计算用户态签名，算法为 HMAC-SHA256(appKey, body)。
func (c *Client) GetUserSignature(body []byte) string {
	return GetUserSignature(body, c.AppKeyProvider)
}

// GetPaySig 计算 pay_sig，算法为 HMAC-SHA256(appKey, uri+"&"+body)。
func GetPaySig(uri string, body []byte, appKey string) string {
	msg := uri + "&" + string(body)
	return hmacHex(appKey, msg)
}

// GetUserSignature 计算用户态签名，算法为 HMAC-SHA256(appKey, body)。
func GetUserSignature(body []byte, appKey string) string {
	return hmacHex(appKey, string(body))
}

func hmacHex(key, msg string) string {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
