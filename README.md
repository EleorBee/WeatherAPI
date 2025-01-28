# WeatherAPI
Sample solution for the [Simple weather API](https://roadmap.sh/projects/weather-api-wrapper-service) challenge from [roadmap.sh](https://roadmap.sh/projects/task-tracker)

## How to run
1. Clone the repository
  ```bash
  git clone https://github.com/EleorBee/WeatherAPI.git
  cd WeatherAPI
  ```
2. Install Go dependencies
  ```bash
  go mod download
  ```
3. Configure Redis and set environment value
```
REDIS_ADDRES   YOUR ADDRES
REDIS_PASSWORD YOUR PASSWORD
REDIS_USER     YOUR USER NAME
API_KEY        YOUR [API KEY](https://www.visualcrossing.com/account) FROM [VISUAL CROSSING](https://www.visualcrossing.com/)
```
4.Run the application
```bash
go run main.go
```
## Usage
Send a GET request to the following endpoint
```
http://localhost:8080/GetWeather/<city_name>
```
