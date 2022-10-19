package logging

import (
	"github.com/google/uuid"
)

func GenerateTransactionID() string {
	ID := uuid.New()
	return ID.String()

}
