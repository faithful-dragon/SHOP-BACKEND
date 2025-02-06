package helper

import (
	database "SHOP-BACKEND/DATABASE"
	utils "SHOP-BACKEND/UTILS"
	"runtime"

	"github.com/kataras/iris/v12"
)

func Onit() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores
	utils.SetCodeMap()                   // set code map
	utils.NewLogger()
	database.ConnectDB()
}

func ServerUp(ctx iris.Context) {
	ctx.HTML("Backend Server Is UP")
}

func SetApiName(apiName string, ctx iris.Context) {
	shop_id := ctx.GetHeader("shop_id")
	logprefix := apiName + "_SHOP_ID_" + shop_id + " : "
	ctx.Values().Set("logPrefix", logprefix)
	ctx.Values().Set("apiName", apiName)
	utils.Logger.Info(logprefix + "Request Recieved.")
}
