package request_input

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIDFromPath(r *http.Request) (int, error) {
	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}

	return ID, nil
}
