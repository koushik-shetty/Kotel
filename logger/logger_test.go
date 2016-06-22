package logger_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koushik-shetty/Kotel/logger"
)

func readFileContents(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

func removeFile(t *testing.T, file string) {
	err := os.Remove(file)
	assert.NoError(t, err, "failed to remove file")
}

func TestNewLoggerCreatesAFile(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.FILE_NAME
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NotNil(t, log, "expected logger to be created")
	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)
	_, err = os.Stat(filename)

	println("err:", err)
	assert.True(t, !os.IsNotExist(err), "Expected file to be created")
	removeFile(t, filename)
}

func TestLoggerWritesEntries(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.FILE_NAME
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)

	log.ErrorF("TestLog: %s", "testdata")

	fileContents, err := readFileContents(filename)
	assert.NoError(t, err, "Expected to read created file")

	println("contents:", fileContents)
	assert.Contains(t, fileContents, "TestLog: testdata", "Expected file to contain log entry")
	removeFile(t, filename)
}

func TestLoggerWritesEntriesOfAppropriateLevel(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.FILE_NAME
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)

	log.InfoF("TestLog: %s", "InfoLog")
	log.ErrorF("TestLog: %s", "ErrorLog")

	fileContents, err := readFileContents(filename)
	assert.NoError(t, err, "Expected to read created file")

	scanner := bufio.NewScanner(bytes.NewBuffer([]byte(fileContents)))
	if scanner.Scan() {
		line := scanner.Text()
		println("text:", line[:4])
		assert.Equal(t, line[:4], "INFO")
	}
	if scanner.Scan() {
		assert.Equal(t, scanner.Text()[:4], "ERRO")
	}
	assert.Contains(t, fileContents, "TestLog: InfoLog", "Expected file to contain log entry")
	assert.Contains(t, fileContents, "TestLog: ErrorLog", "Expected file to contain log entry")
}
