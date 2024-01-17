package routers

import (
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/config"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Load()
}

func MakeRouter(
	config *config.Config,
) *fiber.App {
	cfg := fiber.Config{
		AppName:       "rinha-go by @leorcvargas",
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	}

	r := fiber.New(cfg)

	return r
}
