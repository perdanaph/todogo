package service

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/perdanaph/todoApiGo/internal/todo/model"
	"github.com/perdanaph/todoApiGo/pkg/erru"
)

type CreateParams struct {
	Name        string       `valid:"required"`
	Description string       `valid:"required"`
	Status      model.Status `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (interface{}, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, erru.ErrArgument{Wrapped: err}
	}
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {

		return 0, err
	}
	defer tx.Rollback()

	entity := model.ToDo{
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedOn:   time.Now().UTC(),
	}

	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	return entity, err
}
