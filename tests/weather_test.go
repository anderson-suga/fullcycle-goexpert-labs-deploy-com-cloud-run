package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/dto"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/handler"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/usecase"
)

// --- 1. MOCKS (Simulam a Infraestrutura) ---

type MockCepRepository struct {
	OnGetCity func(cep string) (string, error)
}

func (m *MockCepRepository) GetCity(cep string) (string, error) {
	return m.OnGetCity(cep)
}

type MockWeatherRepository struct {
	OnGetTemp func(city string) (float64, error)
}

func (m *MockWeatherRepository) GetTempCelsius(city string) (float64, error) {
	return m.OnGetTemp(city)
}

// --- 2. TESTES ---

func TestWeatherHandler_EndToEnd(t *testing.T) {
	// Cenário 1: Sucesso (Happy Path)
	t.Run("should return 200 and temps when CEP is valid", func(t *testing.T) {
		// Setup Mocks
		mockCep := &MockCepRepository{
			OnGetCity: func(cep string) (string, error) {
				return "Sao Paulo", nil
			},
		}
		mockWeather := &MockWeatherRepository{
			OnGetTemp: func(city string) (float64, error) {
				return 20.0, nil // Temperatura simulada
			},
		}

		// Setup da Arquitetura
		uc := usecase.NewGetWeatherUseCase(mockWeather, mockCep)
		h := handler.NewWeatherHandler(uc)

		// Request
		req, _ := http.NewRequest("GET", "/weather?cep=01001000", nil)
		rr := httptest.NewRecorder()

		// Execução
		handlerFunc := http.HandlerFunc(h.Handle)
		handlerFunc.ServeHTTP(rr, req)

		// Validações
		if rr.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rr.Code)
		}

		var resp dto.WeatherResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.TempC != 20.0 {
			t.Errorf("expected tempC 20.0, got %f", resp.TempC)
		}
		if resp.TempF != 68.0 { // 20 * 1.8 + 32
			t.Errorf("expected tempF 68.0, got %f", resp.TempF)
		}
		if resp.TempK != 293.0 { // 20 + 273
			t.Errorf("expected tempK 293.0, got %f", resp.TempK)
		}
	})

	// Cenário 2: CEP Inválido (Formato)
	t.Run("should return 422 when CEP format is invalid", func(t *testing.T) {
		// Mocks não devem ser chamados, mas precisamos instanciar
		mockCep := &MockCepRepository{}
		mockWeather := &MockWeatherRepository{}

		uc := usecase.NewGetWeatherUseCase(mockWeather, mockCep)
		h := handler.NewWeatherHandler(uc)

		req, _ := http.NewRequest("GET", "/weather?cep=123", nil)
		rr := httptest.NewRecorder()

		handlerFunc := http.HandlerFunc(h.Handle)
		handlerFunc.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status 422, got %d", rr.Code)
		}
		if rr.Body.String() != "invalid zipcode\n" {
			t.Errorf("expected 'invalid zipcode', got %s", rr.Body.String())
		}
	})

	// Cenário 3: CEP Não Encontrado
	t.Run("should return 404 when CEP is not found", func(t *testing.T) {
		mockCep := &MockCepRepository{
			OnGetCity: func(cep string) (string, error) {
				return "", errors.New("can not find zipcode")
			},
		}
		// Weather não deve ser chamado se CEP falhar
		mockWeather := &MockWeatherRepository{}

		uc := usecase.NewGetWeatherUseCase(mockWeather, mockCep)
		h := handler.NewWeatherHandler(uc)

		req, _ := http.NewRequest("GET", "/weather?cep=99999999", nil)
		rr := httptest.NewRecorder()

		handlerFunc := http.HandlerFunc(h.Handle)
		handlerFunc.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", rr.Code)
		}
		if rr.Body.String() != "can not find zipcode\n" {
			t.Errorf("expected 'can not find zipcode', got %s", rr.Body.String())
		}
	})
}