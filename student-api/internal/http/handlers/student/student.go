package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kishanghosh090/api/internal/types"
	"github.com/kishanghosh090/api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteResponse(w, http.StatusBadRequest, response.GeneralError(errors.New("request body cannot be empty")))

			return
		}
		if err != nil {
			response.WriteResponse(w, http.StatusBadRequest, response.GeneralError(errors.New("invalid request body")))
			return
		}

		// validate request body

		err = validator.New().Struct(student)
		if err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteResponse(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteResponse(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
	}
}
