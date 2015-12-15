package app

import (
	c "Kotel/config"
	"log"
	"os"
)

var logger *log.Logger = nil

//func Write(p []byte) (n int, err error)

func NewFileLogger(lc *c.LoggerConfig) (*log.Logger, error) {
	if logger == nil {
		fileName := lc.Dir() + "/" + lc.FileName()

		os.Mkdir(lc.Dir(), os.ModeDir)
		stream, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}

		log := log.New(stream, "", log.Ldate|log.Ltime|log.Llongfile)
		logger = log
		return log, nil
	}
	return logger, nil

}

func NewConsoleLogger() *log.Logger {
	log := log.New(os.Stdout, ": ", log.Ldate|log.Ltime|log.Llongfile)
	return log
}
