package inventory

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"fmt"
)

type ServerConfig struct {
	InventoryRoot string `yaml:"inventory_root"`
}

func GetConfig() (*ServerConfig, error) {
	yamlRaw, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %s", err)
	}
	serverConfig := &ServerConfig{}
	err = yaml.Unmarshal(yamlRaw, serverConfig)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling config file: %s", err)
	}
	return serverConfig, nil
}
