package file

import (
	"io/ioutil"
	"os"
)

func WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func ReadFile(filename string) ([]byte, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	return ioutil.ReadFile(filename)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
