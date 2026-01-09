package main

import (
	"log"
	"net/http"

	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/config"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/handler"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/infra/viacep"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/infra/weatherapi"
	"github.com/anderson-suga/fullcycle-goexpert-labs-deploy-com-cloud-run/internal/usecase"
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// 2. Set up infrastructure (dependencies)
	cepRepo := viacep.NewClient()
	weatherRepo := weatherapi.NewClient(cfg.WeatherAPIKey)

	// 3. Set up use case
	weatherUseCase := usecase.NewGetWeatherUseCase(weatherRepo, cepRepo)

	// 4. Set up handler
	weatherHandler := handler.NewWeatherHandler(weatherUseCase)

	// 5. Start server
	http.HandleFunc("/weather", weatherHandler.Handle)
	
	log.Printf("Server running on port %s", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, nil); err != nil {
		log.Fatal(err)
	}
}