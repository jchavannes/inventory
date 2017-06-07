package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jchavannes/inventory/inventory"
	"errors"
	"fmt"
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
			inventory.RunServer()
			return nil
		},
	}

	readCmd = &cobra.Command{
		Use:   "read [value]",
		Short: "Read inventory value",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Requires a value to read.")
			}
			value, err := inventory.ReadValue(args[0])
			if err != nil {
				return fmt.Errorf("Error reading value: %s", err)
			}
			fmt.Printf("Value: %s\n", value)
			return nil
		},
	}
)

func Execute() error {
	inventoryCmd.AddCommand(serverCmd)
	inventoryCmd.AddCommand(readCmd)
	return inventoryCmd.Execute()
}
