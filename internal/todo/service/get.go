package service

import (
	"context"
	"errors"

	"github.com/perdanaph/todoApiGo/internal/todo/model"
	"github.com/perdanaph/todoApiGo/pkg/db"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s Service) Get(ctx context.Context, id int) (model.ToDo, error) {
	todo, err := s.repo.Find(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.ToDo{}, erru.ErrArgument{Wrapped: errors.New("todo not found")}
	default:
		return model.ToDo{}, err
	}
	return todo, nil
}
