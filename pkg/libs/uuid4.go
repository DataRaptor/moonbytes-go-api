package libs

import (
	"github.com/google/uuid"
)

func GenerateNewUUID4() string {
	uuid := uuid.New()
	return uuid.String()
}
