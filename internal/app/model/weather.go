package model

type Weather struct {
	Days Days `json:"days"`
}

type Days []struct {
	Date        string  `json:"datetime"`
	Cloudcover  float32 `json:"cloudcover"`
	Description string  `json:"description"`
	Humidity    float32 `json:"humidity"`
	Precip      float32 `json:"precip"`
	Snow        float32 `json:"snow"`
}
