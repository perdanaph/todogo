package service

import (
	"context"
	"errors"

	"github.com/perdanaph/todoApiGo/internal/todo/model"
	"github.com/perdanaph/todoApiGo/pkg/db"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s Service) FindAll(ctx context.Context) ([]model.ToDo, error) {
	todos, err := s.repo.GetAll(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []model.ToDo{}, erru.ErrArgument{Wrapped: errors.New("todo not found")}
	default:
		return []model.ToDo{}, err
	}
	return todos, nil
}
