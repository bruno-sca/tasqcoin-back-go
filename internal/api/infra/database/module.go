package database

import (
	"github.com/bruno-sca/tasqcoin-back-go/internal/api/infra/database/usersdb"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPostgresDatabase),
	usersdb.Module,
)
