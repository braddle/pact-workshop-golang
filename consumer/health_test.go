package consumer_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	consumer "pact-consumer-demo"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestHealthCheck(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "HealthChecker",
		Provider: "DemoHealth",
		LogLevel: "NONE",
	}
	defer pact.Teardown()

	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		Given("The service is up and running").
		UponReceiving("A GET request for the services health").
		WithRequest(
			dsl.Request{
				Method: "GET",
				Path:   dsl.String("/health"),
				Headers: dsl.MapMatcher{
					"Accept":       dsl.String("application/json"),
					"X-Request-Id": dsl.String(requestID),
				},
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusOK,
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.String("application/json"),
				},
				Body: dsl.Match(consumer.HealthResponse{}),
			},
		)

	test := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		hc := consumer.NewHealthChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		hr := hc.Check(ctx)

		assert.Equal(t, "OK", hr.Status)
		assert.Equal(t, int64(36), hr.Integer)
		assert.Equal(t, 12.34, hr.Float)
		assert.False(t, hr.Boolean)

		return nil
	}

	assert.NoError(t, pact.Verify(test))
}
