package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/perdanaph/todoApiGo/internal/todo/model"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s service) FindById() http.HandlerFunc {
	type responseTodo struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}

	type response struct {
		Status  int          `json:"status"`
		Message string       `json:"message"`
		Data    responseTodo `json:"data"`
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

		data, err := s.toDoService.Get(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		s.respond(w, response{
			Status:  200,
			Message: "Success get data",
			Data: responseTodo{
				ID:          data.ID,
				Name:        data.Name,
				Description: data.Description,
				Status:      data.Status,
				CreatedOn:   data.CreatedOn,
				UpdatedOn:   data.UpdatedOn,
			},
		}, http.StatusOK)
	}
}
