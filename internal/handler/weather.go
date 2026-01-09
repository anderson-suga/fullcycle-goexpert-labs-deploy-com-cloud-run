package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/usecase"
)

type WeatherHandler struct {
	UseCase *usecase.GetWeatherUseCase
}

func NewWeatherHandler(uc *usecase.GetWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{UseCase: uc}
}

func (h *WeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cep := r.URL.Query().Get("cep")
	
	resp, err := h.UseCase.Execute(cep)

	if err != nil {
		// Basic error handling mapping
		if err.Error() == "invalid zipcode" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if err.Error() == "can not find zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}