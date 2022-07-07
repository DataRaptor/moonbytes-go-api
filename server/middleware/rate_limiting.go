package middleware

import (
	"gradient-api/server/db"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func RateLimitMiddleware(rateStr string) gin.HandlerFunc {

	rate, err := limiter.NewRateFromFormatted(rateStr)
	if err != nil {
		panic(err)
	}

	store := db.MakeInMemoryStore()

	middleware := mgin.NewMiddleware(
		limiter.New(store, rate))

	return middleware

}
