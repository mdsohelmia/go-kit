package main

import (
	"os"

	"github.com/mdsohelmia/go-kit/bootstrap"
	"github.com/mdsohelmia/go-kit/cmd/migrate"
	"github.com/mdsohelmia/go-kit/config"
	"github.com/mdsohelmia/go-kit/internal/pkg/database"
	"github.com/mdsohelmia/go-kit/internal/pkg/router"
)

func main() {
	app := bootstrap.NewApp()
	// register commands here
	app.AddCommands(migrate.NewMigrateCommand())
	// register modules here
	app.Module(
		config.Module,
		router.Module,
		database.Module,
	)
	if err := app.Execute(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
