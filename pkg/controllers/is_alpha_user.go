package controllers

import (
	"context"
	"gradient-api/pkg/db"
	"gradient-api/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type IsAlphaUserBody struct {
	IsAlphaUser bool `json:"is_alpha_user"`
}

func IsAlphaUserHandler(c *gin.Context) {

	publicKey := c.Param("public_key")
	// TODO: Check that the Wallet Address is on curve.

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

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
	defer cancel()

	var body IsAlphaUserBody

	if len(existing_alpha_users) > 1 {
		panic("Critical Error Alpha Users Have Non-Unqiue Wallet Addresses")
	}

	if len(existing_alpha_users) == 1 {
		body = IsAlphaUserBody{
			IsAlphaUser: true,
		}
	} else {
		body = IsAlphaUserBody{
			IsAlphaUser: false,
		}
	}

	c.JSON(200, body)

}
