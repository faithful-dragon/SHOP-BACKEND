package structs

type ErrorResponse struct {
	ErrorSubResponse ErrorSubResponse `json:"rsp"`
}

type ErrorSubResponse struct {
	ErrorCode       int    `json:"code"`
	ErrorMsg        string `json:"msg"`
	ErrorDescrition string `json:"description"`
}

type AllPartnerDetailsResponse struct {
	AllPartnerDetailsSubResponse AllPartnerDetailsSubResponse `json:"rsp"`
}

type AllPartnerDetailsSubResponse struct {
	PartnerDetails []PartnerDetails `json:"partner_details"`
}

type UniquePartnerDetailsResponse struct {
	AllPartnerDetailsSubResponse UniquePartnerDetailsSubResponse `json:"rsp"`
}

type UniquePartnerDetailsSubResponse struct {
	PartnerDetails PartnerDetails `json:"partner_details"`
}
