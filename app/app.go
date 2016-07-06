package app

import (
	"net/http"
)

func InitApp() (*AppConfig, error) {
	flags := ParseFlags()

	config, err := NewAppConfig(flags)
	return config, err
}

func StartApplication(config *AppConfig) error {
	return http.ListenAndServe(config.GetPort(), nil)
}
