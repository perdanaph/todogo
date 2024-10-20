package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/perdanaph/todoApiGo/internal/todo/model"
	"github.com/perdanaph/todoApiGo/pkg/db"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Find(ctx context.Context, id int) (model.ToDo, error) {
	entity := model.ToDo{}
	query := fmt.Sprintf(
		"SELECT * FROM todos WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}
