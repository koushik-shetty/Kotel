package app

import (
	log "github.com/koushik-shetty/Kotel/logger"
)

type Context struct {
	config AppConfig
	logger log.Logger
}
