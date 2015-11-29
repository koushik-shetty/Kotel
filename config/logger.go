package config

import (
	con "kotel/constants"
)

type LoggerConfig struct {
	fileName string
	dir      string
}

func NewLoggerConfig(dir, file string) *LoggerConfig {
	return &LoggerConfig{
		fileName: file,
		dir:      dir,
	}
}

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		con.LogFileName,
		con.LogDir,
	}
}

func (lc *LoggerConfig) FileName() string {
	return lc.fileName
}

func (lc *LoggerConfig) Dir() string {
	return lc.dir
}
