package controller

import (
	helper "SHOP-BACKEND/HELPER"
	service "SHOP-BACKEND/SERVICE"
	utils "SHOP-BACKEND/UTILS"
	validator "SHOP-BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func AddShopOwner(ctx iris.Context) {
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody := utils.ReadReqBody(ctx)

	utils.Logger.Info(logPrefix, headers, qparams, reqBody)

	errorIdentifier, err := validator.ValidateData(headers, qparams, reqBody)
	rspBody, rspCode := helper.CheckError(errorIdentifier, err)

	var rsp interface{}

	if rspCode != utils.StatusOK {
		rsp = helper.CreateErrorResponse(rspBody)

	} else {
		pid := qparams["pid"]
		shop_id := headers["Shop_id"]
		utils.Logger.Info(logPrefix, pid, shop_id)
		rsp = service.GetUniquePartner(shop_id.(int), pid.(int))
	}
	ctx.JSON(rsp)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
