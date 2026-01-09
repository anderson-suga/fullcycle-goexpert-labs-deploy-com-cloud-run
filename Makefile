.PHONY: test run build clean

# Target to run only the integration tests in the 'tests' folder
test:
	@echo "Running integration tests..."
	go test -v ./tests

# Target to run the application locally (loads .env automatically via code)
run:
	@echo "Running application..."
	go run cmd/server/main.go

# Target to build the binary
build:
	@echo "Building binary..."
	go build -o bin/server cmd/server/main.go

# Target to clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/

# Target to deploy to Google Cloud Run with Free Tier settings
deploy:
	@echo "Deploying to Google Cloud Run (Free Tier Settings)..."
	$(eval API_KEY=$(shell grep WEATHER_API_KEY .env | cut -d '=' -f2))
	gcloud run deploy cep-weather \
		--source . \
		--platform managed \
		--region us-central1 \
		--allow-unauthenticated \
		--memory 512Mi \
		--cpu 1 \
		--min-instances 0 \
		--max-instances 5 \
		--set-env-vars WEATHER_API_KEY=$(API_KEY)