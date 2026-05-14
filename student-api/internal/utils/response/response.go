package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
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
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, "field is required")
		case "email":
			errMsgs = append(errMsgs, "field must be a valid email")
		case "min":
			errMsgs = append(errMsgs, "field must be at least "+err.Param())
		}
	}
	errMsgs = append(errMsgs, "validation failed")

	return Response{
		Status: StatusError,
		Error:  "validation failed: " + strings.Join(errMsgs, ", "),
	}
}
