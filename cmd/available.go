package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo/plugins"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := plugins.Commands{
			{Name: caddyGenCmd.Use, BuffaloCommand: "generate", Description: caddyGenCmd.Short},
			{Name: caddyDevCmd.Use, BuffaloCommand: "dev", Description: caddyDevCmd.Short},
		}
		return json.NewEncoder(os.Stdout).Encode(plugs)
	},
}

func init() {
	RootCmd.AddCommand(availableCmd)
}
