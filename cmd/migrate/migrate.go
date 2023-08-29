package migrate

import (
	"fmt"

	"github.com/mdsohelmia/go-kit/internal/pkg/command"
	"github.com/spf13/cobra"
)

type MigrateCommand struct {
	term string
}

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

func (m *MigrateCommand) Setup(cmd *cobra.Command) {
}

func (m *MigrateCommand) Run(cmd *cobra.Command, args []string) command.CommandRun {
	return func() {
		fmt.Println("run migrate command")
	}
}
