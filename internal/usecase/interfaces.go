package usecase

// WeatherRepository defines the contract for fetching temperature
type WeatherRepository interface {
	GetTempCelsius(city string) (float64, error)
}

// CepRepository defines the contract for fetching location
type CepRepository interface {
	GetCity(cep string) (string, error)
}