package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"

	"gradient-api/pkg/controllers"
	"gradient-api/server/middleware"
)

func CreateRouter() *gin.Engine {

	router := gin.Default()

	router.ForwardedByClientIP = true

	router.ForwardedByClientIP = true

	metrics := ginprometheus.NewPrometheus("gin")
	metrics.Use(router)

	router.Use(middleware.CORSMiddleware())

	router.Any(
		"users/get_alpha_user/:public_key",
		middleware.RateLimitMiddleware("25-M"),
		controllers.GetAlphaUserHandler)

	router.Any(
		"users/create_alpha_user/:public_key",
		middleware.RateLimitMiddleware("3-H"),
		controllers.CreateAlphaUserHandler)

	router.Any(
		"users/is_alpha_user/:public_key",
		middleware.RateLimitMiddleware("25-M"),
		controllers.IsAlphaUserHandler)

	router.Any(
		"data/*proxyPath",
		middleware.RateLimitMiddleware("5-S"),
		controllers.GargantuanAPIReverseProxyHandler)

	return router

}
