package command

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// NewAddCommand creates a new add command
func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add <augend> <addend>",
		Short: "Add two numbers",
		Long:  "Get the sum of two numbers",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			augend, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("invalid augend '%s'", args[0])
			}

			addend, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("invalid addend '%s'", args[1])
			}

			fmt.Println(augend + addend)
		},
	}
}
