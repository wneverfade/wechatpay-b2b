package types

// BatchCreateRetailRequest 预录入门店信息请求参数（Body 部分）。
type BatchCreateRetailRequest struct {
	RetailInfoList []RetailInfo `json:"retail_info_list"`
}

// RetailInfo 门店信息。
type RetailInfo struct {
	MobilePhone        string   `json:"mobile_phone"`
	RetailName         string   `json:"retail_name"`
	RetailType         string   `json:"retail_type,omitempty"`
	SubRetailType      string   `json:"sub_retail_type,omitempty"`
	AddressProvince    string   `json:"address_province"`
	AddressCity        string   `json:"address_city"`
	AddressRegion      string   `json:"address_region"`
	AddressStreet      string   `json:"address_street"`
	RegistrationNumber string   `json:"registration_number,omitempty"`
	BizName            string   `json:"biz_name,omitempty"`
	CorporationName    string   `json:"corporation_name,omitempty"`
	Latitude           float64  `json:"latitude,omitempty"`
	Longitude          float64  `json:"longitude,omitempty"`
	BusinessType       []string `json:"business_type,omitempty"`
	OtherBusinessType  string   `json:"other_business_type,omitempty"`
}

// BatchCreateRetailResponse 预录入门店信息返回参数。
type BatchCreateRetailResponse struct {
	ErrCode           int                           `json:"errcode"`
	ErrMsg            string                        `json:"errmsg"`
	NumSuccess        int                           `json:"num_success,omitempty"`
	NumFailure        int                           `json:"num_failure,omitempty"`
	FailureRecordList []BatchCreateRetailFailRecord `json:"failure_record_list,omitempty"`
}

// BatchCreateRetailFailRecord 单条导入失败记录。
type BatchCreateRetailFailRecord struct {
	MobilePhone        string `json:"mobile_phone"`
	RegistrationNumber string `json:"registration_number,omitempty"`
	FailureCode        int    `json:"failure_code"`
}
