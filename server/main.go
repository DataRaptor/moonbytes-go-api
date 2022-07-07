package main

import (
	pkgDb "gradient-api/pkg/db"
	"gradient-api/server/logger"
	"gradient-api/server/routes"
)

func main() {

	port := ":8080"

	logger.Init()
	pkgDb.ConnectToMongo()
	router := routes.CreateRouter()
	logger.Info("🧙🏻‍♂️ [gradient-api] started on port: " + port)
	router.Run(port)
}
