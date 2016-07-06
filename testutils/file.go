package testutils

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TemporaryFile(name string) string {
	return path.Join(os.TempDir(), name)
}

func CreateTempFile(t *testing.T, name string) *os.File {
	fileName := TemporaryFile(name)

	file, err := os.Create(fileName)
	assert.NoError(t, err, "Expected no error while creating file")

	return file
}

func ReadFileContents(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

func RemoveFile(t *testing.T, name string) {
	err := os.Remove(name)
	assert.NoError(t, err, "Expected no error while removing file")
}
