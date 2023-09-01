package config

import "errors"

var (
	// ErrConfigNotFound is returned when config file is not found
	ErrConfigNotFound = errors.New("config file not found")
	// ErrConfigInvalid is returned when config file is invalid
	ErrConfigInvalid = errors.New("config file is invalid")
	// ErrConfigUnmarshal is returned when config file cannot be unmarshalled
	ErrConfigUnmarshal = errors.New("config file cannot be unmarshalled")
)
