package inventory

import (
	"fmt"
)

func RunServer() error {
	config, err := GetConfig()
	if err != nil {
		return fmt.Errorf("Error getting config file: %s", err)
	}

	fmt.Printf("Inventory root: %s\n", config.InventoryRoot)
	return nil
}
