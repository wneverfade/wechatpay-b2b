package example

import (
	"github.com/enjoy322/wechatpay-b2b/client"
)

// NewClientExample 展示如何初始化 SDK Client（用于 service.NewBalanceService / service.NewRetailService 等）。
// 说明：
// - access_token / appKey 都保存在 client 内，业务侧只需定时刷新并更新到同一个 client 上。
// - 若需要对接多个商户号（mchid），建议按商户号维度维护不同的 client 实例。
func NewClientExample(accessToken, appKey string) (*client.Client, error) {
	return client.NewClient(client.Options{
		BaseURL:        "https://api.weixin.qq.com",
		TokenProvider:  accessToken,
		AppKeyProvider: appKey,
		HTTPClient:     nil,
	})
}
