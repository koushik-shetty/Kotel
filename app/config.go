package app

import (
	"encoding/json"
	"io/ioutil"
)

//config file json format
type ConfigJSON struct {
	Port string `json:"port"`
}

type AppConfig struct {
	port  string
	flags *Flags
}

func (config *AppConfig) GetPort() string {
	return config.port
}

func (config *AppConfig) GetFlags() *Flags {
	return config.flags
}

//genereate a new AppConfig
func NewAppConfig(flags *Flags) (*AppConfig, error) {
	file := flags.ConfigFile
	var config *AppConfig
	if file == "" {
		config = DefaultConfig(flags)
	} else {
		c, err := readConfigFile(file)
		if err != nil {
			return nil, err
		}
		config = c
	}

	config.flags = flags
	return config, nil
}

func DefaultConfig(flags *Flags) *AppConfig {
	port := "666"

	return &AppConfig{
		port:  port,
		flags: flags,
	}
}

func readConfigFile(path string) (*AppConfig, error) {
	if path == "" {
		return DefaultConfig, nil
	}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configJSON := &ConfigJSON{}
	err = json.Unmarshal(configFile, configJSON)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		port: configJSON.Port,
	}, nil
}
