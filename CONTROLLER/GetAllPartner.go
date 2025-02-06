package controller

import (
	helper "SHOP-BACKEND/HELPER"
	service "SHOP-BACKEND/SERVICE"
	utils "SHOP-BACKEND/UTILS"
	validator "SHOP-BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GetAllPartner(ctx iris.Context) {

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody := utils.ReadReqBody(ctx)

	errorIdentifier, err := validator.ValidateData(headers, qparams, reqBody)
	rspBody, rspCode := helper.CheckError(errorIdentifier, err)

	var rsp interface{}

	if rspCode != utils.StatusOK {
		rsp = helper.CreateErrorResponse(rspBody)

	} else {
		shop_id := qparams["shop_id"]
		rsp = service.GetAllPartner(shop_id.(int))
	}
	w := ctx.ResponseWriter()
	w.StatusCode()
	ctx.JSON(rsp)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
