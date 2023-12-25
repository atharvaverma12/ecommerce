package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/atharvaverma12/ecommerce/models"
)

func HashPassword (password string) string{

}

func VerifyPassword (userPassword string, givenPassword string) (bool,string){

}

func Signup() gin.HandlerFunc {
	return func(c* gin.Context){
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()
		var user models.User
		if err:= c.BlindJson(&user); err!=nil{
			c.JSON{http.StatusBadRequest, gin.H{"error":err.Error()}}
			return
		}
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": validationErr})
			return 
		}

		count,err := UserCollection.CountDocuments(ctx,bson.M{"email": user.Email})
		if err!=nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error" : err))
			return
		}

		if count>0{
			c.JSON(http.StatusBadRequest,gin.H{"error":"user already exists"})
		}

		conut, err := UserCollection.CountDocuments(ctx,bson.M{"phone": user.Phone})

		defer cancel()

		if err!=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error": err))
			return
		}

	}
}

func Login() gin.HandlerFunc{

}

func ProductViewerAdmin() gin.HandlerFunc{

}

func searchProduct() gin.HandlerFunc{

}

func SearchProductByQuery() gin.HandlerFunc{

}