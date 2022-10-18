package logging

import (
	log "github.com/sirupsen/logrus"
)

func Log(operationID, layer, function, severity, message string) {
	fields := log.Fields{
		"operation_id": operationID,
		"layer":        layer,
		"function":     function,
	}
	logWithSeverity(fields, severity, message)
}

func logWithSeverity(fields log.Fields, severity, message string) {
	switch severity {
	case INFO:
		log.WithFields(fields).Info(message)
	case WARNING:
		log.WithFields(fields).Warn(message)
	case ALERT:
		log.WithFields(fields).Error(message)
	default:
		log.WithFields(fields).Info(message)
	}
}
