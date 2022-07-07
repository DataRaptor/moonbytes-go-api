package controllers

import (
	"context"
	"fmt"
	"gradient-api/config"
	"gradient-api/pkg/db"
	"gradient-api/pkg/models"
	"gradient-api/server/logger"
	"log"
	"time"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type BadAuthorizationData struct {
	Message string `json:"message"`
}

func GargantuanAPIReverseProxyHandler(c *gin.Context) {

	authorizations, found := c.Request.Header["Authorization"]

	apiKey := authorizations[0]

	fmt.Println(found)

	if !found {
		c.JSON(400, BadAuthorizationData{
			Message: "No Authorization",
		})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"api_key": apiKey}
	cursor, cursorError := db.AlphaUsersCollection.Find(ctx, filter)

	if cursorError != nil {
		logger.Error("Could not Query Mongo for Alpha User")
	}
	// defer cursor.Close(ctx)

	var existing_alpha_users []models.AlphaUser
	err := cursor.All(ctx, &existing_alpha_users)
	if err != nil {
		logger.Error("Could Not Decode Existing Alpha Users")
	}
	// defer cancel()

	fmt.Println(existing_alpha_users)

	if len(existing_alpha_users) != 1 {
		c.JSON(400, BadAuthorizationData{
			Message: "Bad Authorization",
		})
	} else {

		alpha_user := existing_alpha_users[0]
		result, err := db.AlphaUsersCollection.UpdateOne(
			ctx,
			bson.M{"_id": alpha_user.ID},
			bson.D{
				{"$set", bson.D{{"requests_count", alpha_user.RequestsCount + 1}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", result)

		gargantuanAPIUrl := config.GargantuanAPIUrl

		remote, err := url.Parse(gargantuanAPIUrl)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}

}
