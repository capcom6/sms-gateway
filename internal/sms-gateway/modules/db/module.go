package db

import (
	"github.com/jaevor/go-nanoid"
	"go.uber.org/fx"
)

type IDGen func() string

var Module = fx.Module(
	"db",
	fx.Provide(func() (IDGen, error) {
		return nanoid.Standard(21)
	}),
)
