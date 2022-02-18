package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status  string  `json:"Status" pact:"example=OK"`
	Integer int64   `json:"integer" pact:"example=36"`
	Float   float64 `json:"float" pact:"example=12.34"`
	Boolean bool    `json:"boolean" pact:"example=false"`
}

type HealthChecker struct {
	host string
}

func (c HealthChecker) Check(ctx context.Context) HealthResponse {
	url := fmt.Sprintf("%s/health", c.host)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("X-Request-Id", fmt.Sprintf("%v", ctx.Value("request-id")))

	client := http.Client{}
	resp, _ := client.Do(r)

	hr := HealthResponse{}

	json.NewDecoder(resp.Body).Decode(&hr)

	return hr
}

func NewHealthChecker(host string) HealthChecker {
	return HealthChecker{host}
}
