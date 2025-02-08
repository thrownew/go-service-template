package wof

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "wof",
		Short: "Prints 'wof!'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🐶 WOF WOF!")
		},
	}
}
