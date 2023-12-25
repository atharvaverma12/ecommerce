package routes

import(
	"github.com/atharvaverma12/ecommerce/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin, *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login",controllers.Login())
	incomingRoutes.GET("users/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("users/productview",controllers.SearchProduct())
	incomingRoutes.GET("users/search",controllers.SearchProductByQuery())
}

//routes -> controllers -> database