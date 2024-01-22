package model

import "github.com/google/uuid"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Balance  int
	Avatar   string
}

func NewUser(name, email, password string) *User {
	return &User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}
