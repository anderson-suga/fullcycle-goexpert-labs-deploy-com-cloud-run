package dto

// ViaCEPInput represents the raw response from ViaCEP API
type ViaCEPInput struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

// WeatherAPIInput represents the raw response from WeatherAPI
type WeatherAPIInput struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}