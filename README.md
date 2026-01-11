# Full Cycle - Go Expert - Labs - Deploy with Google Cloud Run

This project demonstrates a microservice that receives a Brazilian zipcode (CEP), looks up the location, and returns the current weather conditions (Celsius, Fahrenheit, and Kelvin).

This project was designed to be deployed on Google Cloud Run.

## 🚀 Live Demo

You can access the running application here:
🔗 **[Live URL](https://cep-weather-573316072326.us-central1.run.app/weather?cep=80250070)**

## 📋 Features

- **CEP Validation:** Validates if the input is a valid 8-digit Brazilian CEP.
- **Location Lookup:** Uses **ViaCEP** API to find the city based on the CEP.
- **Weather Lookup:** Uses **WeatherAPI** to fetch current temperature.
- **Unit Conversion:** Automatically converts Celsius to Fahrenheit and Kelvin.
- **Containerized:** Docker & Docker Compose ready.
- **Cloud Native:** Optimized for Google Cloud Run (Serverless).

## 🏗️ Architecture

This project follows **Clean Architecture** principles to ensure decoupling and testability.

````text
cep-weather/
├── cmd/server/       # Application entry point
├── internal/
│   ├── config/       # Environment variables management
│   ├── entity/       # Domain entities
│   ├── usecase/      # Business logic (Pure Go)
│   ├── infra/        # External implementations (HTTP Clients)
│   ├── handler/      # HTTP Transport layer
│   └── dto/          # Data Transfer Objects
├── tests/            # Integration/E2E tests
└── Makefile          # Automation scripts


## ⚙️ Configuration

You need a free API key from WeatherAPI to run this project.

1. Copy the example file:

```bash
cp .env.example .env
````

2. Edit `.env` and add your key:

```text
WEATHER_API_KEY=your_api_key_here
PORT=8080
```

## 🛠️ How to Run

### Option 1: Using Makefile (Recommended)

Run the application locally:

```bash
make run
```

The server will start at http://localhost:8080

## 🧪 Testing

The project includes integration tests that mock external APIs.

To run the tests:

```bash
make test
```

## 📡 API Reference

**Request:** GET /weather?cep={cep}

**Example:** GET /weather?cep=80250070

#### Success Response (200 OK):

```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.0
}
```

#### Error Responses:

- 422 Unprocessable Entity: Invalid CEP format.
  - invalid zipcode
- 404 Not Found: CEP not found.
  - can not find zipcode
- 500 Internal Server Error: External API failure.

## 🛠️ Makefile Commands

The repository includes a `Makefile` with the most used tasks. Quick reference:

- `make test`: Run integration tests in the `tests/` folder.
- `make run`: Run the application locally (loads `.env` via code).
- `make build`: Build the application binary at `bin/server`.
- `make clean`: Remove build artifacts (`bin/`).
- `make deploy`: Deploy the service to Google Cloud Run using `gcloud run deploy`.
  - This target reads `WEATHER_API_KEY` from your local `.env` and sets it as an environment variable in the Cloud Run service.

</br>

> This is a postgraduate degree challenge from [Full Cycle - Go Expert](https://goexpert.fullcycle.com.br/pos-goexpert/)
