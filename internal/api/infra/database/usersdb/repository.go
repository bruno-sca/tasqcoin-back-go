package usersdb

import (
	"context"

	"github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersSQLRepository struct {
	db *pgxpool.Pool
}

func (r *UsersSQLRepository) Create(user *model.User) error {
	_, err := r.db.Exec(context.Background(),
		"INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *UsersSQLRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(context.Background(),
		"SELECT id, name, email, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func NewUsersSQLRepository(db *pgxpool.Pool) *UsersSQLRepository {
	return &UsersSQLRepository{
		db: db,
	}
}
