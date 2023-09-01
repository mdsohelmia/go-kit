package config

import "go.uber.org/fx"

// package bootstrap
var Module = fx.Module("config", fx.Options(
	fx.Provide(NewConfig),
))
