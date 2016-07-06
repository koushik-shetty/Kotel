package constants

const (
	DevModeFlag    = "dev"
	DevModeFlagMsg = "-dev=true for development mode"
	DefDevMode     = false

	ConfigFilePath = "config"
	DefConfigFile  = ""
	ConfigFileMsg  = "-config=<filepath>, json config for the application"

	LogDirFlag    = "logdir"
	LogDirFlagMsg = "directory to store the logs. Should be relative to the current directory"
	LogDir        = "logs"
	LogFileName   = "kotel.logs.txt"

	AppName = "Western Wall"
)
