package app

func InitApp() AppConfig {
	return AppConfig{
		Flags:  getFlags(),
		Config: getConfig(),
	}
}

func StartApplication(config *AppConfig) error {
	http.ListenAndServer(config.GetConfig().Port)
}
