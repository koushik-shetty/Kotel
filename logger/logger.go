package logger

import (
	"io"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
)

const (
	FILE_NAME = "Kotel.log"
)

//TODO: move logging into seperate repo so as to share the logging across different projects
type Loggable interface {
	InfoF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
	FatalF(format string, args ...interface{})
	PrintF(format string, args ...interface{})
}

//adapter type around logrus. CLients are removed from the external dependency
type Logger struct {
	textLogger *logrus.Logger
}

//TODO: implement log rotation
//returns a new logger creating a new file in the process
func NewLogger(level string, logDir string, logFile string) (*Logger, error) {
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	filename := fileName(logFile)
	file, err := createLogFile(logDir, filename)
	if err != nil {
		return nil, err
	}

	logger := &logrus.Logger{
		Level:     parsedLevel,
		Out:       file,
		Formatter: new(logrus.TextFormatter),
	}
	return &Logger{
		textLogger: logger,
	}, nil
}

func (l *Logger) InfoF(format string, args ...interface{}) {
	l.textLogger.Infof(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.textLogger.Error(args...)
}

func (l *Logger) ErrorF(format string, args ...interface{}) {
	l.textLogger.Errorf(format, args...)
}

func (l *Logger) FatalF(format string, args ...interface{}) {
	l.textLogger.Fatalf(format, args...)
}

func (l *Logger) PrintF(format string, args ...interface{}) {
	l.textLogger.Printf(format, args)
}

//internal helper functions
func createLogFile(logDir string, logFile string) (io.Writer, error) {
	if ok, err := dirExists(logDir); !ok {
		return nil, err
	}

	file, err := os.Create(path.Join(logDir, logFile))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func fileName(file string) string {
	if file == "" {
		return FILE_NAME
	}
	return file
}
