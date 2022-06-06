package health

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseWithPactExamples struct {
	// pact tags are used to configure example response values to be used in tests
	Status  string  `json:"Status" pact:"example=NOT OK"`
	Integer int64   `json:"integer" pact:"example=36"`
	Float   float64 `json:"float" pact:"example=12.34"`
	Boolean bool    `json:"boolean" pact:"example=false"`
}

type ResponseWithoutPactExamples struct {
	Status  string  `json:"Status"`
	Integer int64   `json:"integer"`
	Float   float64 `json:"float"`
	Boolean bool    `json:"boolean"`
}

type HealthChecker struct {
	host string
}

// Check makes a HTTP call to the Producers healthcheck endpoint /health
func (c HealthChecker) Check(ctx context.Context) ResponseWithPactExamples {
	// Creating the URL to request provide host's /health endpoint
	url := fmt.Sprintf("%s/health", c.host)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	// Creating the require Accept header to inform the provider we require a JSON response
	r.Header.Add("Accept", "application/json")
	// Forwarding the required X-Request-Id
	r.Header.Add("X-Request-Id", fmt.Sprintf("%v", ctx.Value("request-id")))

	client := http.Client{}

	// Make the HTTP request. The error from Do() is not handled because you should not be using Pact to test errors where
	// we fail to connect to the API. This functionality should be cover by separate Unit Tests.
	resp, _ := client.Do(r)

	hr := ResponseWithPactExamples{}

	// Decoding the JSON response from the API. The error from Decode() is not handled because you should not be using
	//Pact to test errors where we fail to connect to the API. This functionality should be cover by separate Unit Tests.
	json.NewDecoder(resp.Body).Decode(&hr)

	return hr
}

// NewChecker creates an instance of a Healthcheck please provide the hostname for the service providing the /health
// endpoint (Example: http://api.testing.com)
func NewChecker(hostname string) HealthChecker {
	return HealthChecker{hostname}
}
