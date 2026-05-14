package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

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
		println(student.Email)

		response.WriteResponse(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
	}
}
