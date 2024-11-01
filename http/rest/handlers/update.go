package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/perdanaph/todoApiGo/internal/todo/model"
	toDoService "github.com/perdanaph/todoApiGo/internal/todo/service"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name        *string       `json:"name"`
		Description *string       `json:"description"`
		Status      *model.Status `json:"status"`
	}

	type response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		req := request{}

		err = s.decode(r, &req)

		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.toDoService.Update(r.Context(), toDoService.UpdateParams{
			ID:          id,
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			Status:  http.StatusOK,
			Message: "Success update data",
		}, http.StatusOK)
	}
}
