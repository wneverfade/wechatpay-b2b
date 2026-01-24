# wechatpay-b2b

微信支付 B2B（门店助手）Golang SDK。

## 安装

```bash
go get github.com/enjoy322/wechatpay-b2b
```

## 项目定位

- 面向小程序 B2B 门店助手的"支付/退款/提现/合单支付"等能力，提供后端 SDK 封装（签名、请求结构体、调用入口等）。
- 小程序侧支付入口为 `wx.requestCommonPayment`，本 SDK 负责在服务端准备 `signData`、`paySig` 等参数并对接相关服务端接口。

## 官方文档

- 文档：<https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html>
- 获取小程序下所有商户的信息：<https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html#_8-%E8%8E%B7%E5%8F%96%E5%B0%8F%E7%A8%8B%E5%BA%8F%E4%B8%8B%E6%89%80%E6%9C%89%E5%95%86%E6%88%B7%E7%9A%84%E4%BF%A1%E6%81%AF>

## 支持的接口

### 商户服务 (MerchantService)

| 方法 | 功能 | URI |
|-----|-----|-----|
| `GetMerchantInfo` | 获取小程序下所有商户的信息 | `/retail/B2b/getmchinfo` |
| `GetMerchantAppKey` | 查询商户的 appKey | `/retail/B2b/getappkey` |
| `GetBalance` | 查询账户余额 | `/retail/B2b/getmchbalance` |
| `Withdraw` | 发起提现 | `/retail/B2b/withdraw` |
| `QueryWithdraw` | 查询提现状态 | `/retail/B2b/querywithdraw` |

### 订单服务 (OrderService)

| 方法 | 功能 | URI |
|-----|-----|-----|
| `CloseOrder` | 关闭订单 | `/retail/B2b/closeb2border` |
| `GetOrder` | 查询订单 | `/retail/B2b/getorder` |
| `CreateRefund` | 发起退款 | `/retail/B2b/createrefund` |
| `GetRefund` | 查询退款 | `/retail/B2b/getrefund` |
| `BuildPaymentParams` | 生成单订单支付参数 | `requestCommonPayment` |
| `BuildCombinedPaymentParams` | 生成合单支付参数 | `requestCommonPayment` |

`BuildPaymentParams` / `BuildCombinedPaymentParams` 仅生成小程序支付参数，不发起 HTTP 请求。

### 分账服务 (ProfitService)

| 方法 | 功能 | URI |
|-----|-----|-----|
| `ProfitSharing` | 请求分账 | `/retail/B2b/profitsharing` |
| `QueryProfitSharing` | 查询分账订单 | `/retail/B2b/queryprofitsharing` |
| `ProfitSharingFinish` | 分账完结 | `/retail/B2b/profitsharingfinish` |
| `ProfitSharingReturn` | 分账回退 | `/retail/B2b/profitsharingreturn` |
| `QueryProfitSharingReturn` | 查询分账回退 | `/retail/B2b/queryprofitsharingreturn` |

### 门店服务 (RetailService)

| 方法 | 功能 | URI |
|-----|-----|-----|
| `BatchCreateRetail` | 预录入门店信息 | `/wxa/business/batchcreateretail` |

### 通知解析

| 函数 | 功能 |
|-----|-----|
| `model.ParsePaymentNotify` | 解析支付通知 |
| `model.ParseRefundNotify` | 解析退款通知 |

## 使用说明

### Client 配置（access_token / appKey 传参）

本 SDK 以 `client.Client` 作为共享调用上下文：

- `access_token`：保存在 `client.Client` 内，业务侧定时刷新并更新即可。
- `appKey`：不保存在 `client.Client` 内，调用需要 `pay_sig` 的服务时传入（建议按商户号维度维护 `map[mchid]appKey`）。
- `session_key`：不保存在 `client.Client` 内，调用 `OrderService.BuildPaymentParams` / `BuildCombinedPaymentParams` 时传入，用于计算 `signature`。

因此调用 `service.MerchantService.GetBalance` 等方法时，需要显式传入 `appKey`，但 `access_token` 只需在 `client.Client` 内保持最新即可。

建议直接替换整个 `client.Client` 实例。

### 示例

```go
package main

import (
    "context"
    "fmt"

    "github.com/enjoy322/wechatpay-b2b/client"
    "github.com/enjoy322/wechatpay-b2b/service"
    "github.com/enjoy322/wechatpay-b2b/types"
)

func main() {
    // 初始化客户端
    c, err := client.NewClient(client.Options{
        BaseURL:     "https://api.weixin.qq.com",
        AccessToken: "your_access_token",
    })
    if err != nil {
        panic(err)
    }
    appKey := "your_app_key"

    // 创建服务
    orderSvc := service.NewOrderService(c)

    // 查询订单
    resp, err := orderSvc.GetOrder(context.Background(), types.GetOrderRequest{
        Mchid:      "1230000109",
        OutTradeNo: "your_out_trade_no",
    }, appKey)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Order: %+v\n", resp)
}
```
