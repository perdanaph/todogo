package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(handler.MiddlewareLogger())
	apiRouter.HandleFunc("/todo", handler.Create()).Methods(http.MethodPost)
	apiRouter.HandleFunc("/todos", handler.GetAllTodos()).Methods(http.MethodGet)
	apiRouter.HandleFunc("/todo/{id}", handler.FindById()).Methods(http.MethodGet)
	apiRouter.HandleFunc("/todo/{id}", handler.Update()).Methods(http.MethodPut)
	apiRouter.HandleFunc("/todo/{id}", handler.Delete()).Methods(http.MethodDelete)
}
