package inventory

import (
	"fmt"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"strings"
)

func ReadValue(keyRaw string) (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", fmt.Errorf("Error getting config file: %s", err)
	}
	appFileRaw, err := ioutil.ReadFile(config.InventoryRoot + "/app.sls")
	if err != nil {
		return "", fmt.Errorf("Error reading app file: %s", err)
	}
	var yamlData map[string]interface{}
	err = yaml.Unmarshal(appFileRaw, &yamlData)
	if err != nil {
		return "", fmt.Errorf("Error parsing yaml: %s", err)
	}

	keys := strings.Split(keyRaw, ":")
	for i, key := range keys {
		for k, v := range yamlData {
			if k != key {
				continue
			}
			switch v := v.(type) {
			case map[interface{}]interface{}:
				yamlData = convertMapIItoSI(v)
			case string:
				if i == len(keys) - 1 {
					return v, nil
				}
			}
			break
		}
	}
	return "", fmt.Errorf("Unable to find key: %s", keyRaw)
}

func convertMapIItoSI(input map[interface{}]interface{}) map[string]interface{} {
	returnMap := make(map[string]interface{})
	for k, v := range input {
		switch k := k.(type) {
		case string:
			returnMap[k] = v
		}
	}
	return returnMap
}
