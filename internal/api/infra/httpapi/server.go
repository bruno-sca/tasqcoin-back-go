package httpapi

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/config"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func NewServer(
	lifecycle fx.Lifecycle,
	router *fiber.App,
	config *config.Config,
) *fasthttp.Server {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Info("Starting the server...")

				addr := fmt.Sprintf(":%s", config.Server.Port)
				if err := router.Listen(addr); err != nil {
					log.Fatalf("Error starting the server: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping the server...")
			return router.ShutdownWithContext(ctx)
		},
	})

	return router.Server()
}
