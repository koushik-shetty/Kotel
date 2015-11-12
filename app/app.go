package app

import (
	"net/http"
)

func InitApp() AppConfig {
	return AppConfig{
		Flags:  getFlags(),
		Config: getConfig(),
	}
}

func StartApplication(config *AppConfig) error {
	http.ListenAndServe(config.GetConfig().Port, nil)
	return nil
}
