package WeatherApi

import (
	"WeatherAPI/internal/app/caching"
	"WeatherAPI/internal/app/endpoint"
	"WeatherAPI/internal/app/service"
	"WeatherAPI/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/khaaleoo/gin-rate-limiter/core"
	"golang.org/x/time/rate"
	"os"
)

type App struct {
	router   *gin.Engine
	caching  *caching.Caching
	endpoint *endpoint.Endpoint
	service  *service.Service

	rateLimiter gin.HandlerFunc

	port string
}

func New() *App {

	cfg := config.MustLoadConfig()

	router := gin.Default()
	service := service.New(cfg.Endpoint, os.Getenv("API_KEY"))
	caching := caching.New(service)
	endpoint := endpoint.New(caching, service)

	rateLimitedOption := core.RateLimiterOption{
		Limit: rate.Limit(cfg.Limit),
		Burst: cfg.MaxRequest,
		Len:   cfg.ResetLimit,
	}
	rateLimitedMiddleware := core.RequireRateLimiter(core.RateLimiter{
		RateLimiterType: core.IPRateLimiter,
		Key:             "GetWeatherLimited",
		Option:          rateLimitedOption,
	})

	return &App{
		caching:     caching,
		endpoint:    endpoint,
		service:     service,
		router:      router,
		port:        cfg.Port,
		rateLimiter: rateLimitedMiddleware,
	}
}

func (app *App) Run() {

	app.router.GET("GetWeather/:city", app.rateLimiter, app.endpoint.Weather)

	err := app.router.Run(":8080")

	if err != nil {
		fmt.Println("error working server:", err)
	}
}
