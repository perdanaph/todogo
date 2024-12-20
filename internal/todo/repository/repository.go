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

func (r Repository) GetAll(ctx context.Context) ([]model.ToDo, error) {
	var todos []model.ToDo
	query := `
		SELECT id, name, description, status, created_on, updated_on, deleted_on 
		FROM todos 
		WHERE deleted_on IS NULL
	`
	err := r.Db.SelectContext(ctx, &todos, query)
	if err != nil {
		return nil, db.HandleError(err)
	}
	return todos, nil
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

func (r Repository) Update(ctx context.Context, entity model.ToDo) error {
	query := `UPDATE todo
                SET name = :name, 
                    description = :description, 
                    status = :status, 
                    created_on = :created_on, 
                    updated_on = :updated_on, 
                    deleted_on = :deleted_on
                WHERE id = :id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return db.HandleError(err)
}

func (r Repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM todo WHERE id = $1`
	_, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return db.HandleError(err)
	}
	return nil
}
