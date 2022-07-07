package db

import (
	limiter "github.com/ulule/limiter/v3"

	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func MakeInMemoryStore() limiter.Store {
	store := memory.NewStore()
	return store
}
