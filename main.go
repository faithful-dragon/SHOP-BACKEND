package main

import (
	controller "SHOP-BACKEND/CONTROLLER"
	helper "SHOP-BACKEND/HELPER"

	"github.com/kataras/iris/v12"
)

func main() {

	helper.Onit()

	app := iris.New()

	// to get server is up
	app.Get("/shop/ping", func(ctx iris.Context) {
		helper.ServerUp(ctx)
	})

	// api to add new shop owner
	app.Post("shop/addShowOwner", func(ctx iris.Context) {
		helper.SetApiName("ADD_SHOP_OWNER", ctx)
		controller.AddShopOwner(ctx)
	})

	// api to get unique partner
	app.Get("/shop/uniquePartner", func(ctx iris.Context) {
		helper.SetApiName("GET_UNIQUE_PARTNER", ctx)
		controller.GetUniquePartner(ctx)
	})

	// api to get all partner
	app.Get("/shop/all_partner", func(ctx iris.Context) {
		helper.SetApiName("GET_ALL_PARTNER", ctx)
		controller.GetAllPartner(ctx)
	})

	// Start the server on port 8000
	err := app.Listen(":8000")
	if err != nil {
		app.Logger().Fatal(err)
	}
}
