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

func (r Repository) Create(ctx context.Context, entity *model.ToDo) error {
	query := `INSERT INTO todos (name, description, status, created_on, updated_on)
                VALUES (:name, :description, :status, :created_on, :updated_on) RETURNING id;`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}
	return db.HandleError(err)
}
