package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
func IsJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}

func IsFileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}
	return info.IsDir(), nil
}

//dwd/wd
