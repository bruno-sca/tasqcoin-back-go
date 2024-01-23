package users

import (
	"errors"

	"github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	PASSWORD_HASH_COST = 14
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORD_HASH_COST)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us *UserService) Create(user *model.User) error {
	userFound, err := us.repo.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if userFound != nil {
		return errors.New("user already exists")
	}

	user.Password, err = HashPassword(user.Password)

	if err != nil {
		return err
	}

	us.repo.Create(user)

	return nil
}
