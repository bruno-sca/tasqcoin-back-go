package controllers

import (
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/model"
	users "github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/service"
	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	service users.IUserService
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (uc *UsersController) Create(c *fiber.Ctx) error {
	var body CreateUserRequest

	if err := c.BodyParser(&body); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	err := uc.service.Create(model.NewUser(body.Name, body.Email, body.Password))

	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(fiber.Map{"message": "user created"})
}

func NewUsersController(service users.IUserService) *UsersController {
	return &UsersController{
		service: service,
	}
}
