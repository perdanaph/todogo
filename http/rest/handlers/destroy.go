package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s service) Delete() http.HandlerFunc {
	type response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must be filled in path"),
			}, 0)
			return
		}

		error := s.toDoService.Destroy(r.Context(), id)
		if error != nil {
			s.respond(w, error, 0)
			return
		}

		s.respond(w, response{
			Status:  200,
			Message: "Success get data",
		}, http.StatusNoContent)
	}
}
