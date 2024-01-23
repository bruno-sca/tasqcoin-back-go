package usersdb

import (
	users "github.com/bruno-sca/tasqcoin-back-go/internal/api/domain/users/service"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewUsersSQLRepository,
		fx.As(new(users.Repository)),
	),
)
