package users

import (
	model "github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/model"
)

type IUserService interface {
	Create(user *model.User) error
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}
