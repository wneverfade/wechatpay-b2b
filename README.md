# wechatpay-b2b

微信支付 B2B（门店助手）Golang SDK。

## 项目定位

- 面向小程序 B2B 门店助手的“支付/退款/提现/合单支付”等能力，提供后端 SDK 封装（签名、请求结构体、调用入口等）。
- 小程序侧支付入口为 `wx.requestCommonPayment`，本 SDK 负责在服务端准备 `signData`、`paySig` 等参数并对接相关服务端接口。

## 官方文档

- 文档：<https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/B2b_store_assistant.html#_3-3-%E6%94%AF%E4%BB%98%E4%B8%8E%E9%80%80%E6%AC%BE>

## 使用说明

### Client 配置（access_token / appKey）

本 SDK 以 `client.Client` 作为共享调用上下文：

- `access_token`：保存在 `client.Client.TokenProvider`，业务侧定时刷新并更新即可。
- `appKey`：保存在 `client.Client.AppKeyProvider`，用于服务端接口计算 `pay_sig`（建议按商户号维度维护不同的 `client.Client` 实例）。

因此调用 `service.BalanceService.GetBalance` 等方法时，不需要每次显式传入 `access_token` / `appKey`，只需确保 `client.Client` 内的相关字段是最新的。

建议直接替换整个 `client.Client` 实例。
