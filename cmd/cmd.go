package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jchavannes/inventory/server"
)

var (
	inventoryCmd = &cobra.Command{
		Use:   "inventory",
		Short: "Run the inventory application",
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run inventory server",
		RunE: func(c *cobra.Command, args []string) error {
			server.Run()
			return nil
		},
	}
)

func Execute() error {
	inventoryCmd.AddCommand(serverCmd)
	return inventoryCmd.Execute()
}
