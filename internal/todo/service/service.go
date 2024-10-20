package service

import "github.com/perdanaph/todoApiGo/internal/todo/repository"

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
