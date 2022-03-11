package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status  string  `json:"Status" pact:"example=OK"` // pact tags are used to configure example response values to be used in tests
	Integer int64   `json:"integer" pact:"example=36"`
	Float   float64 `json:"float" pact:"example=12.34"`
	Boolean bool    `json:"boolean" pact:"example=false"`
}

type HealthChecker struct {
	host string
}

// Check makes a HTTP call to the Producers healthcheck endpoint /health
func (c HealthChecker) Check(ctx context.Context) HealthResponse {
	url := fmt.Sprintf("%s/health", c.host) // Creating the URL to request provide host's /health endpoint
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Accept", "application/json")                               // Creating the require Accept header to inform the provider we require a JSON response
	r.Header.Add("X-Request-Id", fmt.Sprintf("%v", ctx.Value("request-id"))) // Forwarding the required X-Request-Id

	client := http.Client{}

	// Make the HTTP request. The error from Do() is not handled because you should not be using Pact to test errors where
	// we fail to connect to the API. This functionality should be cover by separate Unit Tests.
	resp, _ := client.Do(r)

	hr := HealthResponse{}

	// Decoding the JSON response from the API. The error from Decode() is not handled because you should not be using
	//Pact to test errors where we fail to connect to the API. This functionality should be cover by separate Unit Tests.
	json.NewDecoder(resp.Body).Decode(&hr)

	return hr
}

// NewHealthChecker creates an instance of a Healthcheck please provide the hostname for the service providing the /health
// endpoint (Example: http://api.testing.com)
func NewHealthChecker(hostname string) HealthChecker {
	return HealthChecker{hostname}
}
