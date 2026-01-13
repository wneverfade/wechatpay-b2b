package types

// RegisterMerchantRequest 商户号进件请求参数。
type RegisterMerchantRequest struct {
	IDDocTypeNum         uint32          `json:"id_doc_type_num"`
	IDCardInfo           *IDCardInfo     `json:"id_card_info,omitempty"`
	IDDocInfo            *IDDocInfo      `json:"id_doc_info,omitempty"`
	AccountInfo          AccountInfo     `json:"account_info"`
	ContactInfo          ContactInfo     `json:"contact_info"`
	BusinessLicense      BusinessLicense `json:"business_license"`
	MerchantShortname    string          `json:"merchant_shortname"`
	OrganizationType     uint32          `json:"organization_type"`
	Qualification        *Qualification  `json:"qualification,omitempty"`
	BusinessAdditionDesc string          `json:"business_addition_desc,omitempty"`
	BusinessAdditionPics string          `json:"business_addition_pics,omitempty"`
	OpenType             uint32          `json:"open_type"`
	ExtRegisterInfo      ExtRegisterInfo `json:"ext_register_info"`
	ClientIP             string          `json:"client_ip"`
}

// IDCardInfo 经营者/法人身份证信息。
type IDCardInfo struct {
	IDCardCopy           string `json:"id_card_copy"`
	IDCardNational       string `json:"id_card_national"`
	IDCardName           string `json:"id_card_name"`
	IDCardNumber         string `json:"id_card_number"`
	IDCardValidTime      string `json:"id_card_valid_time"`
	IDCardAddress        string `json:"id_card_address"`
	IDCardValidTimeBegin string `json:"id_card_valid_time_begin"`
}

// IDDocInfo 经营者/法人其他类型证件信息。
type IDDocInfo struct {
	IDDocName      string `json:"id_doc_name"`
	IDDocNumber    string `json:"id_doc_number"`
	IDDocCopy      string `json:"id_doc_copy"`
	DocPeriodEnd   string `json:"doc_period_end"`
	DocPeriodBegin string `json:"doc_period_begin"`
	IDDocAddress   string `json:"id_doc_address"`
	IDDocCopyBack  string `json:"id_doc_copy_back"`
}

// AccountInfo 结算银行账户信息。
type AccountInfo struct {
	BankAccountType string `json:"bank_account_type"`
	AccountBank     string `json:"account_bank"`
	AccountName     string `json:"account_name"`
	BankAddressCode string `json:"bank_address_code"`
	BankBranchID    string `json:"bank_branch_id,omitempty"`
	BankName        string `json:"bank_name,omitempty"`
	AccountNumber   string `json:"account_number"`
}

// ContactInfo 超级管理员信息。
type ContactInfo struct {
	ContactType                 string `json:"contact_type"`
	ContactName                 string `json:"contact_name"`
	ContactIDDocType            string `json:"contact_id_doc_type,omitempty"`
	ContactIDCardNumber         string `json:"contact_id_card_number,omitempty"`
	ContactIDDocCopy            string `json:"contact_id_doc_copy,omitempty"`
	ContactIDDocCopyBack        string `json:"contact_id_doc_copy_back,omitempty"`
	ContactIDDocPeriodBegin     string `json:"contact_id_doc_period_begin,omitempty"`
	ContactIDDocPeriodEnd       string `json:"contact_id_doc_period_end,omitempty"`
	BusinessAuthorizationLetter string `json:"business_authorization_letter,omitempty"`
	MobilePhone                 string `json:"mobile_phone"`
	ContactEmail                string `json:"contact_email,omitempty"`
}

// BusinessLicense 营业执照信息。
type BusinessLicense struct {
	BusinessLicenseCopy   string `json:"business_license_copy"`
	BusinessLicenseNumber string `json:"business_license_number"`
	MerchantName          string `json:"merchant_name"`
	LegalPerson           string `json:"legal_person"`
	CompanyAddress        string `json:"company_address,omitempty"`
	BusinessTime          string `json:"business_time,omitempty"`
	CertType              string `json:"cert_type,omitempty"`
}

// Qualification 行业特殊资质资料。
type Qualification struct {
	QualificationType string   `json:"qualification_type"`
	Qualifications    []string `json:"qualifications,omitempty"`
}

// ExtRegisterInfo 补充信息。
type ExtRegisterInfo struct {
	DoorHeadFileID            string `json:"door_head_file_id,omitempty"`
	StoreFileID               string `json:"store_file_id,omitempty"`
	OnlinePayFileID           string `json:"online_pay_file_id,omitempty"`
	MerchantScale             string `json:"merchant_scale,omitempty"`
	AuthorizationLetterFileID string `json:"authorization_letter_file_id,omitempty"`
	ContactIDDocAddress       string `json:"contact_id_doc_address,omitempty"`
}

// RegisterMerchantResponse 商户号进件返回参数。
type RegisterMerchantResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	OrderNo string `json:"order_no,omitempty"`
}

// GetMerchantOpenStatusRequest 查询商户号开通状态请求参数。
type GetMerchantOpenStatusRequest struct {
	OutRegistrationID string `json:"out_registration_id,omitempty"`
	PageIndex         uint32 `json:"page_index,omitempty"`
	PageSize          uint32 `json:"page_size,omitempty"`
}

// GetMerchantOpenStatusResponse 查询商户号开通状态返回参数。
type GetMerchantOpenStatusResponse struct {
	ErrCode int                      `json:"errcode"`
	ErrMsg  string                   `json:"errmsg"`
	List    []MerchantOpenStatusItem `json:"list,omitempty"`
	Total   uint32                   `json:"total,omitempty"`
}

// MerchantOpenStatusItem 商户号进件订单状态信息。
type MerchantOpenStatusItem struct {
	Status               uint32                        `json:"status,omitempty"`
	InnerResp            *MerchantOpenStatusInnerResp  `json:"inner_resp,omitempty"`
	WqfRegisterStatement *MerchantWqfRegisterStatement `json:"wqf_register_statement,omitempty"`
	WxPayRate            int32                         `json:"wx_pay_rate,omitempty"`
	WqfCertifiedRate     int32                         `json:"wqf_certified_rate,omitempty"`
	BindSceneStatus      uint32                        `json:"bind_scene_status,omitempty"`
}

// MerchantOpenStatusInnerResp 进件订单内部状态信息。
type MerchantOpenStatusInnerResp struct {
	SubMerchantRegistrationStatus *SubMerchantRegistrationStatus `json:"sub_merchant_registration_status,omitempty"`
}

// SubMerchantRegistrationStatus 申请状态信息。
type SubMerchantRegistrationStatus struct {
	ApplymentState     string                     `json:"applyment_state"`
	ApplymentStateDesc string                     `json:"applyment_state_desc"`
	SignState          string                     `json:"sign_state,omitempty"`
	SignURL            string                     `json:"sign_url,omitempty"`
	SubMchid           string                     `json:"sub_mchid,omitempty"`
	AccountValidation  *MerchantAccountValidation `json:"account_validation,omitempty"`
	AuditDetail        []MerchantAuditDetail      `json:"audit_detail,omitempty"`
	LegalValidationURL string                     `json:"legal_validation_url,omitempty"`
}

// MerchantAccountValidation 汇款账户验证信息。
type MerchantAccountValidation struct {
	AccountName              string `json:"account_name"`
	AccountNo                string `json:"account_no"`
	PayAmount                int64  `json:"pay_amount"`
	DestinationAccountNumber string `json:"destination_account_number"`
	DestinationAccountName   string `json:"destination_account_name"`
	DestinationAccountBank   string `json:"destination_account_bank"`
	City                     string `json:"city"`
	Remark                   string `json:"remark"`
	Deadline                 string `json:"deadline"`
}

// MerchantAuditDetail 驳回原因详情。
type MerchantAuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

// MerchantWqfRegisterStatement 银行转账开通状态。
type MerchantWqfRegisterStatement struct {
	WqfRegisterState     uint32 `json:"wqf_register_state"`
	WqfRegisterStateDesc string `json:"wqf_register_state_desc"`
	RequestNo            string `json:"request_no,omitempty"`
}
