package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	toDoRepo "github.com/perdanaph/todoApiGo/internal/todo/repository"
	toDoService "github.com/perdanaph/todoApiGo/internal/todo/service"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	toDoService toDoService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:      lg,
		toDoService: toDoService.NewService(toDoRepo.NewRepository(db)),
	}
}
