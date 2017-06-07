package server

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

type ServerConfig struct {
	InventoryRoot string `yaml:"inventory_root"`
}

func GetConfig() (*ServerConfig, error) {
	yamlRaw, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	serverConfig := &ServerConfig{}
	err = yaml.Unmarshal(yamlRaw, serverConfig)
	if err != nil {
		return nil, err
	}
	return serverConfig, nil
}
