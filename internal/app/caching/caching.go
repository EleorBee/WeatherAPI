package caching

import (
	"WeatherAPI/internal/app/model"
	"WeatherAPI/internal/app/service"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type Caching struct {
	rdb *redis.Client
	srv *service.Service
}

func New(service *service.Service) *Caching {

	c := &Caching{
		rdb: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Username: os.Getenv("REDIS_USERNAME"),
			DB:       0,
		}),
		srv: service,
	}
	return c
}

func (c *Caching) SaveWeather(key string, weather *model.Weather) {
	ctx := context.Background()

	value, err := json.Marshal(weather)

	timeNow := time.Now()

	data := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)

	expiration := time.Until(data)

	if err = c.rdb.Set(ctx, key, value, expiration).Err(); err != nil {
		log.Println("failed to cache:", key)
	}
}

func (c *Caching) GetWeather(city string) (*model.Weather, error) {
	ctx := context.Background()
	val, err := c.rdb.Get(ctx, city).Bytes()

	var weather *model.Weather

	if err == redis.Nil {

		weather, err = c.srv.Request(city)

		if err != nil {
			return nil, err
		}

		c.SaveWeather(city, weather)

		return weather, nil
	} else {
		err = json.Unmarshal(val, &weather)

		return weather, err
	}
}
