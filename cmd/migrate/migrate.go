package migrate

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mdsohelmia/go-kit/config"
	"github.com/mdsohelmia/go-kit/internal/pkg/command"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
)

const (
	dir = "database/migrations"
)

type MigrateCommand struct{}

var _ command.Command = (*MigrateCommand)(nil)

func NewMigrateCommand() *MigrateCommand {
	return &MigrateCommand{}
}

func (m *MigrateCommand) Aliases() []string {
	return []string{"m"}
}

func (m *MigrateCommand) SuggestFor() []string {
	return []string{}
}

func (m *MigrateCommand) Long() string {
	return ""
}

func (m *MigrateCommand) Example() string {
	return `usage:
migrate up
migrate up-by-one
migrate up-to
migrate down
migrate down-to 20170506082527
migrate status
migrate redo
migrate create
migrate version
`
}

func (m *MigrateCommand) Short() string {
	return "Migrate database schema"
}

func (m *MigrateCommand) Use() string {
	return "migrate"
}

func (m *MigrateCommand) Setup(cmd *cobra.Command) {}

func (m *MigrateCommand) Run(cmd *cobra.Command, args []string) command.CommandRun {
	return func(config *config.Config, database *bun.DB) {
		if len(args) == 0 {
			fmt.Println(m.Example())
			return
		}
		defer database.Close()

		action, args := args[0], args[1:]

		if err := goose.SetDialect("mysql"); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		err := goose.Run(action, database.DB, dir, args...)

		if err != nil {
			fmt.Println(strings.ReplaceAll(err.Error(), `\n`, "\n"))
		}
	}
}

func (m *MigrateCommand) getMigrationPath(cfg *config.Config) string {
	driver := cfg.Database.Backend.Driver
	switch driver {
	case config.DatabaseDriverMysql:
		return cfg.Database.Backend.Mysql.MigrationPath
	case config.DatabaseDriverTidb:
		return cfg.Database.Backend.Tidb.MigrationPath
	default:
		return "database/migrations"
	}
}
