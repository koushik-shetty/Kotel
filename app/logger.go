package app

import (
	c "Kotel/config"
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
)

func NewFileLogger(lc *c.LoggerConfig) (*logrus.Logger, error) {
	fileName := lc.Dir() + "/" + lc.FileName()

	os.Mkdir(lc.Dir(), os.ModeDir)
	stream, err := os.Create(fileName)
	fmt.Printf("create:%v,%v,%v\n", fileName, err, stream)

	if err != nil {
		return nil, err
	}

	logrus.SetOutput(stream)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.ErrorLevel)

	log := logrus.New()
	return log, nil
}
