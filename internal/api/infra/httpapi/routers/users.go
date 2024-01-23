package routers

import (
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
)

type UsersRouter struct {
	controller *controllers.UsersController
}

func (ur *UsersRouter) Load(r *fiber.App) {
	r.Post("/users", ur.controller.Create)
}

func NewUsersRouter(controller *controllers.UsersController) *UsersRouter {
	return &UsersRouter{controller: controller}
}
