package main

import (
	users "github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/service"
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/config"
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/database"
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/httpapi"
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/httpapi/controllers"
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/httpapi/routers"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main() {
	uuid.EnableRandPool()

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	app := fx.New(
		config.Module,
		users.Module,
		controllers.Module,
		routers.Module,
		httpapi.Module,
		database.Module,
		fx.Invoke(func(*fasthttp.Server) {}),
	)

	app.Run()
}
