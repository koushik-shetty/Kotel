package app

import (
	"flag"
	// "log"
)

type AppConfig struct {
	DevMode bool
	LogDir  string
}

func (conf *AppConfig) InitApp() {
	devMode := flag.Bool("dev", false, "-dev=true for development mode")
	logDir := flag.String("logdir", "./logs", "directory to store the logs. Should be relative to the current directory")

	flag.Parse()

	conf.DevMode = *devMode
	conf.LogDir = *logDir

}
