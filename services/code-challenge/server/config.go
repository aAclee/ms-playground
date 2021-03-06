package main

import (
	"encoding/json"
	"os"
)

// Configuration type for code-challenge server configs
type Configuration struct {
	Port int `json:"port"`
}

func getServerConfigs(filename string) (*Configuration, error) {
	configuration := Configuration{}

	err := Parse(filename, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

// Parse parses a json path/filename and fills the configuration
func Parse(filename string, configuration *Configuration) error {
	config, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(config)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}

	return nil
}
