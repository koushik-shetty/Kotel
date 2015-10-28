package app

import (
	"flag"
	// "log"
	"errors"
)

type Flags struct {
	DevMode bool
	LogDir  string
}

type Config struct {
	Port string
}

var DefaultConfig = Config{
	Port: "5004",
}

type AppConfig struct {
	Config Config
	Flags  Flags
}

func (config *AppConfig) GetConfig() Config {
	return config.Config
}

func (config *AppConfig) GetFlags() Flags {
	return config.Flags
}

func getFlags() Flags {

	devMode := flag.Bool(devModeFlag, false, devModeFlagMsg)
	logDir := flag.String(logDirFlag, logDirDefault, logDirFlagMsg)

	flag.Parse()

	return Flags{
		DevMode: *devMode,
		LogDir:  *logDir,
	}
}

func getConfig() Config {
	config, err := readConfigFile("")
	if err != nil {
		return DefaultConfig
	}
	return *config
}

func readConfigFile(path string) (config *Config, err error) {
	return nil, errors.New("Failed to read config file")
}
