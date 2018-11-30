package cli

import (
	"github.com/charliekenney23/drone-test/pkg/cli/command"

	"github.com/spf13/cobra"
)

// Cli is the command line interface
var Cli = cobra.Command{Use: "math"}

func init() {
	Cli.AddCommand(command.NewAddCommand())
}
