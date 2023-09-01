package bootstrap

import (
	"context"
	"fmt"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mdsohelmia/go-kit/internal/pkg/command"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version:          "0.0.1",
	Use:              "go-kit",
	Short:            "go-kit is a toolkit for building microservices in Go",
	Long:             `go-kit is a toolkit for building microservices in Go`,
	TraverseChildren: true,
}

type App struct {
	cmd      *cobra.Command
	commands []*cobra.Command
	options  []fx.Option
}

func NewApp() *App {
	return &App{
		cmd: rootCmd,
	}
}

func (a *App) Module(opts ...fx.Option) {
	a.options = append(a.options, opts...)
}

func (a *App) AddCommands(cmds ...command.Command) {
	for _, cmd := range cmds {
		wrappedCmd := &cobra.Command{
			Use:              cmd.Use(),
			Aliases:          cmd.Aliases(),
			SuggestFor:       cmd.SuggestFor(),
			Short:            cmd.Short(),
			Long:             cmd.Long(),
			Example:          cmd.Example(),
			TraverseChildren: true,
			Run: func(c *cobra.Command, args []string) {
				ctx := context.Background()
				app := fx.New(
					fx.Options(a.options...),
					fx.NopLogger,
					fx.Invoke(cmd.Run(c, args)),
				)
				err := app.Start(ctx)
				if err != nil {
					fmt.Println(err)
				}
				defer app.Stop(ctx)
			},
		}
		cmd.Setup(wrappedCmd)
		a.commands = append(a.commands, wrappedCmd)
	}

}

func (a *App) Execute() error {
	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Aliases:       cc.Bold + cc.Italic,
		CmdShortDescr: cc.HiRed,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDescr:    cc.HiRed,
		FlagsDataType: cc.Italic,
	})

	a.cmd.AddCommand(a.commands...)
	return a.cmd.Execute()
}
