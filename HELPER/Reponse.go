package helper

import (
	structs "SHOP-BACKEND/STRUCTS"
	utils "SHOP-BACKEND/UTILS"
)

func CreateErrorResponse(rsp utils.Codes) structs.ErrorResponse {
	return structs.ErrorResponse{
		ErrorSubResponse: structs.ErrorSubResponse{
			ErrorCode:       rsp.StatusCode,
			ErrorMsg:        rsp.Message,
			ErrorDescrition: rsp.Description,
		},
	}
}

func CreateUniquePartnerResponse(rsp structs.PartnerDetails) structs.UniquePartnerDetailsResponse {
	return structs.UniquePartnerDetailsResponse{
		AllPartnerDetailsSubResponse: structs.UniquePartnerDetailsSubResponse{
			PartnerDetails: rsp,
		},
	}
}

func CreateGetAllPartnerResponse(rsp utils.Codes) []structs.AllPartnerDetailsResponse {
	return []structs.AllPartnerDetailsResponse{
		{
			AllPartnerDetailsSubResponse: structs.AllPartnerDetailsSubResponse{
				PartnerDetails: []structs.PartnerDetails{},
			},
		},
	}
}
