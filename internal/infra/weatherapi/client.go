package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/dto"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) GetTempCelsius(city string) (float64, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", c.apiKey, encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weather API error: %v", resp.Status)
	}

	var wInput dto.WeatherAPIInput
	
	if err := json.NewDecoder(resp.Body).Decode(&wInput); err != nil {
		return 0, err
	}

	return wInput.Current.TempC, nil
}