package handlers

import (
	"net/http"
	"time"

	"github.com/perdanaph/todoApiGo/internal/todo/model"
)

func (s service) GetAllTodos() http.HandlerFunc {

	type responseTodos struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}

	type response struct {
		Status  int             `json:"status"`
		Message string          `json:"message"`
		Data    []responseTodos `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.toDoService.FindAll(r.Context())
		if err != nil {
			s.respond(w, err, 500)
			return
		}

		var todosResponse []responseTodos
		for _, todos := range data {
			todosResponse = append(todosResponse, responseTodos{
				ID:          todos.ID,
				Name:        todos.Name,
				Description: todos.Description,
				Status:      todos.Status,
				CreatedOn:   todos.CreatedOn,
				UpdatedOn:   todos.UpdatedOn,
			})
		}
		s.respond(w, response{
			Status:  http.StatusOK,
			Message: "Success get all data",
			Data:    todosResponse,
		}, http.StatusOK)
	}
}
