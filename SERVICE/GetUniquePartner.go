package service

import (
	database "SHOP-BACKEND/DATABASE"
	helper "SHOP-BACKEND/HELPER"
	structs "SHOP-BACKEND/STRUCTS"
	utils "SHOP-BACKEND/UTILS"
	"database/sql"
)

func GetUniquePartner(shopId int, pid int) interface{} {

	var partnerDetails structs.PartnerDetails
	DB := database.ConnectDB()
	defer DB.Close() // Close database connection after function execution

	ServiceQuery := database.GetUniquePartner() // SQL query string

	err := DB.QueryRow(ServiceQuery, pid).Scan(&partnerDetails.FullName, &partnerDetails.CompanyName, &partnerDetails.PhNo, &partnerDetails.Address, &partnerDetails.RegDate)

	if err != nil {
		utils.Logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return helper.ErrorResponse("400001") // Partner not found
		}
		return helper.ErrorResponse("400003") // Database/Scanning error
	}

	response := helper.CreateUniquePartnerResponse(partnerDetails)
	return response
}
