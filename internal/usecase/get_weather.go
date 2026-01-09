package usecase

import (
	"fmt"
	"regexp"

	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/dto"
)

type GetWeatherUseCase struct {
	WeatherRepo WeatherRepository
	CepRepo     CepRepository
}

func NewGetWeatherUseCase(w WeatherRepository, c CepRepository) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		WeatherRepo: w,
		CepRepo:     c,
	}
}

func (u *GetWeatherUseCase) Execute(cep string) (*dto.WeatherResponse, error) {
	if !isValidCEP(cep) {
		return nil, fmt.Errorf("invalid zipcode")
	}

	city, err := u.CepRepo.GetCity(cep)
	if err != nil {
		return nil, err
	}

	tempC, err := u.WeatherRepo.GetTempCelsius(city)
	if err != nil {
		return nil, err
	}

	return &dto.WeatherResponse{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}, nil
}

func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^[0-9]{8}$`)

	return re.MatchString(cep)
}