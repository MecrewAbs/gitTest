package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const BaseURL = "https://api.jsonbin.io/v3"

type Client struct {
	APIKey string
	Client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: &http.Client{},
	}
}

// CreateBin отправляет новый bin на JSON-BIN.
func (c *Client) CreateBin(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", BaseURL+"/b", strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("API error: %d %s", resp.StatusCode, string(body))
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result["id"], nil
}

// GetBin получает bin по ID.
func (c *Client) GetBin(id string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", BaseURL+"/b/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Master-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %d %s", resp.StatusCode, string(body))
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
