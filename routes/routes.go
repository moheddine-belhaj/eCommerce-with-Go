package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moheddine-belhaj/eCommerce-with-Go/controllers"

)

var filename = "./product.json"


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewAdmin())
	incomingRoutes.POST("/admin/AddProductsFromFile",controllers.AddProductsFromFile(filename))
	incomingRoutes.GET("/users/productview",controllers.SearchProduct())
	incomingRoutes.GET("/admin/all",controllers.GetAll())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuerry())

}	