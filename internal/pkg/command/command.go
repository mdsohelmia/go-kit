package command

import "github.com/spf13/cobra"

type CommandRun any

// Command is the interface that wraps the basic methods of a command.
type Command interface {
	Setup(cmd *cobra.Command)
	Use() string
	Aliases() []string
	SuggestFor() []string
	Long() string
	Example() string
	Short() string
	Run(cmd *cobra.Command, args []string) CommandRun
}
