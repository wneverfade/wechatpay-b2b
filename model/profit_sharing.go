package model

// ProfitReceiverType 分账接收方类型。
type ProfitReceiverType string

const (
	// ProfitReceiverMerchant 商户
	ProfitReceiverMerchant ProfitReceiverType = "MERCHANT"
	// ProfitReceiverIndividual 个人
	ProfitReceiverIndividual ProfitReceiverType = "INDIVIDUAL"
)

// ProfitReceiverRelationType 分账接收方关系类型。
type ProfitReceiverRelationType string

const (
	// RelationServiceProvider 服务商
	RelationServiceProvider ProfitReceiverRelationType = "SERVICE_PROVIDER"
	// RelationStore 门店
	RelationStore ProfitReceiverRelationType = "STORE"
	// RelationEmployee 员工
	RelationEmployee ProfitReceiverRelationType = "EMPLOYEE"
	// RelationOther 其他
	RelationOther ProfitReceiverRelationType = "OTHER"
)

// ProfitReceiver 分账接收方信息。
type ProfitReceiver struct {
	Type      ProfitReceiverType        `json:"type"`                 // 分账接收方类型
	Account   string                    `json:"account"`              // 分账接收方账号
	Amount    int64                     `json:"amount"`               // 分账金额
	Desc      string                    `json:"desc,omitempty"`       // 分账描述
	Relation  ProfitReceiverRelationType `json:"relation"`             // 分账接收方关系
}

// ProfitStatus 分账状态。
type ProfitStatus string

const (
	// ProfitStatusInit 初始化
	ProfitStatusInit ProfitStatus = "PROFIT_SHARING"
	// ProfitStatusFinished 已分账完成
	ProfitStatusFinished ProfitStatus = "PROFIT_FINISHED"
	// ProfitStatusClosed 已关闭
	ProfitStatusClosed ProfitStatus = "PROFIT_CLOSED"
)

// ProfitReturnStatus 分账回退状态。
type ProfitReturnStatus string

const (
	// ProfitReturnPending 待回退
	ProfitReturnPending ProfitReturnStatus = "RETURN_PENDING"
	// ProfitReturnSuccess 回退成功
	ProfitReturnSuccess ProfitReturnStatus = "RETURN_SUCCESS"
	// ProfitReturnFailed 回退失败
	ProfitReturnFailed ProfitReturnStatus = "RETURN_FAILED"
)
