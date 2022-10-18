package logging

import (
	"github.com/google/uuid"
)

func GenerateOperationID() string {
	ID := uuid.New()
	return ID.String()

}
