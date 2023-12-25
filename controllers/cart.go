package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/atharvaverma12/ecommerce/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection * mongo.Collection 
	userCollection * mongo.Collection 
}

func NewApplication(prodCollection,userCollection * mongo.Collection) *Application{
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection
	}
}

func (app * Application)AddToCart() gin.Handler {
	return func (c * gin.Context){

		//check for product id
		productQuery := c.Query("id")
		if productQuery == ""{
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest,errors.New("product is is empty"))
			return
		}

		//check for user id
		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}
		
		//if prod is is compatible
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err!=nil {
			println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx,app.prodCollection,app.userCollection, productID, userQueryID)

		if err!= nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200,"successfully")
	}
}

func (app * Application) RemoveItem() gin.HandlerFunc {
return func (c)
}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
