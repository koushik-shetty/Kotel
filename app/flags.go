package app

import (
	"flag"

	con "github.com/koushik-shetty/Kotel/constants"
)

type Flags struct {
	DevMode    bool
	ConfigFile string
}

func NewFlags(devmode bool, configFile string) *Flags {
	return &Flags{
		DevMode:    devmode,
		ConfigFile: configFile,
	}
}

func ParseFlags() *Flags {
	devMode := flag.Bool(con.DevModeFlag, con.DefDevMode, con.DevModeFlagMsg)
	config := flag.String(con.ConfigFilePath, con.DefConfigFile, con.ConfigFileMsg)

	flag.Parse()

	return &Flags{
		DevMode:    *devMode,
		ConfigFile: *config,
	}
}

func DefaultFlags() *Flags {
	return &Flags{
		DevMode:    con.DefDevMode,
		ConfigFile: con.DefConfigFile,
	}
}
