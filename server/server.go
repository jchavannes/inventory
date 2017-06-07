package server

import "fmt"

func Run() error {
	config, err := GetConfig()
	if err != nil {
		return err
	}
	fmt.Printf("Inventory root: %s\n", config.InventoryRoot)
	return nil
}
