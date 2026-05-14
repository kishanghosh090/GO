package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status string
	Error  string
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func WriteResponse(
	w http.ResponseWriter,
	status int,
	data interface{},
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// ... rest of the function
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		status: StatusError,
		Error:  err.Error(),
	}
}
