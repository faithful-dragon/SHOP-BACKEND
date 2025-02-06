package service

import (
	helper "SHOP-BACKEND/HELPER"
)

func GetAllPartner(shopId int) interface{} {

	// ServiceQuery := database.GetAllPartner()
	// rows, err := database.DB.Query(ServiceQuery, shopId)
	// if err != nil {
	// 	return helper.ErrorResponse("400001")
	// }
	// defer rows.Close()
	return helper.ErrorResponse("400001")
}
