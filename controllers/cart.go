package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/atharvaverma12/ecommerce/database"
	"github.com/atharvaverma12/ecommerce/models"
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

			var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
			defer cancel()

			err = database.RemoveCartItem(ctx,app.prodCollection, app.userCollection, productID, userQueryID)

			if err!=nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}

			c.IndentedJSON(200, "successfully removed item from cart")

}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c * gin.Context){
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type","application/json")
			c.JSON(http.StatusNotFound,gin.H{"error":"invalid id"})
			c.Abort()
			return
		}

		usert_id , _ := primitive.ObjectIDFromHex(user_id)

		//whenever we work with a database we need to define the context first
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledcart models.User
		err := UserCollection.FindOne(ctx,bson.D{primitive.E{Key: "_id", Value: usert_id}}).Decode(&filledcart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		//we get data of user ie match
		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}

		//unwind the data of the  ie unwind
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		
		//start working with the value of the user ie group
		grouping := bson.D{{Key: "group", Value: bson.D{primitive.E{Key:"_id", Value: "$_id"}, {Key:"total", Value: bson.D{primitive.E{Key: "sum", Value:"$usercart.price"}}}}}}
		
		//run the aggregation func
		pointcursor, err := UserCollection.Aggregate(ctx,mongo.Pipeline{filter_match,unwind,grouping})
		if err != nil {
			log.Println(err)
		}
		var listing [] bson.M

		//convert all the values of pointcursor to listing
		if err = pointcursor.All(ctx,&listing); err != nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		//send total to user in JSON format
		for _,json := range listing{
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200,filledcart.UserCart)
		}
		ctx.Done()
	}
}

func (app * Application) BuyFromCart() gin.HandlerFunc {
	return func(c * gin.Context){
		userQueryID := c.Query("id")
		if userQueryID == ""{
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(),100 * time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx,app.userCollection,userQueryID)

		if(err!=nil){
			c.IndentedJSON(http.StatusInternalServerError,err)
			return
		}
		c.IndentedJSON("200","successfully placed the order")
	}
}

func (app * Application) InstantBuy() gin.HandlerFunc {
	return func (c* gin.Context){

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
		
					var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
					defer cancel()

					err = database.InstantBuyerctx,app.prodCollection, app.userCollection, productID, userQueryID)

					if err!=nil{
						c.IndentedJSON(http.StatusInternalServerError, err)
						return
					}
					c.IndentedJSON(200, "successfully bought the item!")
	}
}
