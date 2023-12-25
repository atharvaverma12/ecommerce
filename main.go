package main

import {
	"github.com/atharvaverma12/ecommerce/controllers"
	"github.com/atharvaverma12/ecommerce/routes"
	"github.com/atharvaverma12/ecommerce/middleware"
	"github.com/atharvaverma12/ecommerce/database"
	"github.com/gin-gonic/gin"
}

func main(){
	//start the server
	port := os.Getend("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client,"Products"),database.UserData(database.Client,"Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem",app.RemoveItem())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/instantbuy",app.InstantBuy())
}