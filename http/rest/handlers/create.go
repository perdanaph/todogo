package handlers

import (
	"net/http"

	"github.com/perdanaph/todoApiGo/internal/todo/model"
	toDoService "github.com/perdanaph/todoApiGo/internal/todo/service"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
	}
	type response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(r, &req)
		if err != nil {
			// fmt.Println("masuk sini")
			s.respond(w, err, 0)
			return
		}
		data, err := s.toDoService.Create(r.Context(), toDoService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			// fmt.Println("masuk sini")
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			Status:  http.StatusCreated,
			Message: "Success add data",
			Data:    data,
		}, http.StatusCreated)
	}
}
