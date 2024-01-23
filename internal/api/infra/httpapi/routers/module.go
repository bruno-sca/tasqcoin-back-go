package routers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewUsersRouter,
	MakeRouter,
)
