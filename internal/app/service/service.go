package service

import (
	"WeatherAPI/internal/app/model"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"time"
)

type Service struct {
	client   *resty.Client
	Endpoint string
	ApiKey   string
}

func New(endpoint string, apiKey string) *Service {
	return &Service{
		client:   resty.New(),
		Endpoint: endpoint,
		ApiKey:   apiKey,
	}
}

func (s *Service) Request(location string) (*model.Weather, error) {
	r, err := s.client.R().
		SetQueryParam("key", s.ApiKey).
		Get(s.Endpoint + location + "/" + time.Now().Format(time.DateOnly))

	if err != nil {
		log.Fatal("request bad!:", err.Error())
		return nil, err
	}

	if http.StatusBadRequest == r.StatusCode() {
		return nil, fmt.Errorf("City not found")
	}
	if r.StatusCode() == http.StatusUnauthorized {
		return nil, fmt.Errorf("There is a problem with the API key, account or subscription")
	}
	if r.StatusCode() == http.StatusTooManyRequests {
		return nil, fmt.Errorf("The web server service has exceeded its assigned limits.")
	}
	if r.StatusCode() == http.StatusInternalServerError {
		return nil, fmt.Errorf("A general error has occurred processing the request")
	}

	var res *model.Weather
	err = json.Unmarshal(r.Body(), &res)

	if err != nil {
		log.Fatal("unmarshal error:", err.Error())
		return nil, err
	}

	return res, nil
}
