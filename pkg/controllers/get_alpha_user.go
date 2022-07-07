package controllers

import (
	"context"
	"gradient-api/pkg/db"
	"gradient-api/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAlphaUserHandler(c *gin.Context) {

	publicKey := c.Param("public_key")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var alpha_user models.AlphaUser
	defer cancel()

	filter := bson.M{"public_key": publicKey}
	cur, currErr := db.AlphaUsersCollection.Find(ctx, filter)

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var existing_alpha_users []models.AlphaUser
	err := cur.All(ctx, &existing_alpha_users)
	if err != nil {
		panic(err)
	}

	if len(existing_alpha_users) > 1 {
		panic("Critical Error Alpha Users Have Non-Unqiue Wallet Addresses")
	}

	if len(existing_alpha_users) == 1 {
		alpha_user = existing_alpha_users[0]
		c.JSON(200, alpha_user)
	} else {
		empty_alpha_user := models.AlphaUser{
			Uuid:          "",
			PublicKey:     "",
			ApiKey:        "",
			RequestsCount: 0,
		}
		c.JSON(200, empty_alpha_user)
	}

}
