package service

import (
	"context"
	"errors"

	"github.com/perdanaph/todoApiGo/pkg/db"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s Service) Destroy(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)

	switch {
	case err != nil:
		return nil

	case errors.As(err, &db.ErrObjectNotFound{}):
		return erru.ErrArgument{Wrapped: errors.New("todo not found")}
	default:
		return err
	}
}
