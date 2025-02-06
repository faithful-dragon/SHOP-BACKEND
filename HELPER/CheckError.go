package helper

import (
	utils "SHOP-BACKEND/UTILS"
)

func ErrorResponse(errorCode string) utils.Codes {
	return utils.CodeMap[errorCode]
}

func CheckError(identifier string, err error) (utils.Codes, int) {
	rspBody := ErrorResponse("200001")
	rspCode := utils.StatusOK
	switch identifier {
	case "INVALID_HEADER":
		if err != nil {
			utils.Logger.Error(err.Error())
			rspBody = ErrorResponse("400001")
			rspCode = utils.StatusBadRequest
		}
	case "INVALID_QPARAMS":
		if err != nil {
			utils.Logger.Error(err.Error())
			rspBody = ErrorResponse("400002")
			rspCode = utils.StatusBadRequest
		}
	case "INVALID_BODY":
		if err != nil {
			utils.Logger.Error(err.Error())
			rspBody = ErrorResponse("400003")
			rspCode = utils.StatusBadRequest
		}
	}

	return rspBody, rspCode
}
