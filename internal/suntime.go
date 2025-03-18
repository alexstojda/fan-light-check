package suntime

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const sunriseSunsetUrl = "https://api.sunrisesunset.io/json?lat=45.5031824&lng=-73.5698065&time_format=unix&timezone=UTC"

type Response struct {
	Results Times `json:"results"` 
}

type Times struct {
	Sunrise Timestamp `json:"sunrise"`
	Sunset Timestamp `json:"sunset"`
}

func GetTimes() (*Times, error) {
	resp, err := http.Get(sunriseSunsetUrl)
	if err != nil {
		return nil, fmt.Errorf("performing request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %v", err)
	}

	data := &Response{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling data: %v", err)
	}

	return &data.Results, nil
}
