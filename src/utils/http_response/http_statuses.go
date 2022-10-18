package http_response

import (
	"clean_architecture_in_go/src/utils/statuses"
	"net/http"
)

var HTTPStatuses = map[string]int{
	statuses.SUCCESS:        http.StatusOK,
	statuses.CONFLICT:       http.StatusConflict,
	statuses.INTERNAL_ERROR: http.StatusInternalServerError,
	statuses.INPUT_ERROR:    http.StatusBadRequest,
	statuses.NOT_FOUND:      http.StatusNotFound,
}
