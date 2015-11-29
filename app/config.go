package app

import (
	"errors"
	"flag"
	con "kotel/constants"
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

	devMode := flag.Bool(con.DevModeFlag, false, con.DevModeFlagMsg)
	logDir := flag.String(con.LogDirFlag, con.LogDir, con.LogDirFlagMsg)

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
