package storage

import (
	"encoding/json"
	"os"
)

type Bin struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func SaveToJSON(bins []Bin, filename string) error {
	data, err := json.Marshal(bins)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func LoadFromJSON(fileName string) ([]Bin, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var bins []Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		return nil, err
	}
	return bins, nil
}

//dwdwd
