package config

import (
	"fmt"

	"go.uber.org/fx"
)

type Config struct{}

func NewConfig() *Config {
	fmt.Println("NewConfig")
	return &Config{}
}

// package bootstrap
var Module = fx.Module("config", fx.Options(
	fx.Provide(NewConfig),
))
