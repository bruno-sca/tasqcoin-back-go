package users

import "github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/model"

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
