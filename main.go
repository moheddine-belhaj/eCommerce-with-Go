package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moheddine-belhaj/eCommerce-with-Go/controllers"
	"github.com/moheddine-belhaj/eCommerce-with-Go/database"
	"github.com/moheddine-belhaj/eCommerce-with-Go/middleware"
	"github.com/moheddine-belhaj/eCommerce-with-Go/routes"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8002"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client,"Users"))
	routers := gin.New()

	routers.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentifcation())

	router.GET("/addtocart",app.AddToCart())
	router.GET("reomveitem",app.RomoveItem())
	router.GET("cartcheckout",app.BuyFromCart())
	router.GET("/instanbuy",app.instanBuy())

	log.Fatal((router.run(":" + port)))
}