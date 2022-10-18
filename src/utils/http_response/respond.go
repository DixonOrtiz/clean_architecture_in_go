package http_response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
}

func Respond(w http.ResponseWriter, status string, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatuses[status])

	response := Response{Status: status}

	if data != nil {
		response.Data = data
	}

	if err != nil {
		response.Data = nil
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
