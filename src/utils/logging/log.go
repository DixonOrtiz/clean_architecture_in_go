package logging

import (
	log "github.com/sirupsen/logrus"
)

func NewLog(transactionID, layer, function string) *log.Logger {
	return log.WithFields(log.Fields{
		"transaction_id": transactionID,
		"layer":          layer,
		"function":       function,
	}).Logger
}
