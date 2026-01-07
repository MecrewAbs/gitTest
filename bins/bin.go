package bins

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Bin struct {
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	SecretKey string `json:"secretKey"`
}

func NewBin(name string, private bool) *Bin {
	return &Bin{
		Name:    name,
		Private: private}
}

func (b *Bin) SetPrivate() {
	b.Private = true
	b.SecretKey = generateKey(32)
}

func (b *Bin) IsPrivate() bool {
	return b.Private
}

func (b *Bin) ToJSON() (string, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (b *Bin) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), b)
}

func generateKey(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func (b *Bin) String() string {
	return fmt.Sprintf("name: %s, private: %t, secretKey: %s", b.Name, b.Private, b.SecretKey)
}
