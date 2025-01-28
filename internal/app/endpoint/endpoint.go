package endpoint

import (
	"WeatherAPI/internal/app/caching"
	"WeatherAPI/internal/app/model"
	"WeatherAPI/internal/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Endpoint struct {
	c *caching.Caching
	s *service.Service
}

func New(c *caching.Caching, s *service.Service) *Endpoint {
	return &Endpoint{c, s}

}

func (e *Endpoint) Weather(c *gin.Context) {
	city := c.Param("city")
	var weather *model.Weather
	var err error

	weather, err = e.c.GetWeather(city)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": weather})
	}
}
