package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetCity(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error calling ViaCEP")
	}

	var rawMap map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawMap); err != nil {
		return "", err
	}

	if _, ok := rawMap["erro"]; ok {
		return "", fmt.Errorf("can not find zipcode")
	}

	localidade, ok := rawMap["localidade"].(string)
	if !ok {
		return "", fmt.Errorf("error parsing location")
	}

	return localidade, nil
}