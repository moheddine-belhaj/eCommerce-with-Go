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
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client,"Users"))
	router := gin.New()

	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart",app.AddToCart())
	router.GET("/reomveitem",app.RemoveItem())	
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/instanbuy",app.InstantBuy())
	router.POST("/addadress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/listcart", controllers.GetItemFromCart())

	log.Fatal((router.Run(":" + port)))
}