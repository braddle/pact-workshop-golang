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
	// Instantiating Pact
	//
	// NOTE: If you are running multiple tests again the same Provider create one instance dls.Pact that is used by all tests.
	pact := &dsl.Pact{
		Consumer: "Health Checker Client", // The name of the consumer using the API. In this cade out Health Checker Client
		Provider: "Demo Health Endpoint",  // The name of the Provider we testing against
		LogLevel: "NONE",
	}

	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		Given("The service is up and running").                 // Providing the expectations for the Provider to setup
		UponReceiving("A GET request for the services health"). // Describing the request that will be made
		WithRequest(
			// Configuring the request you expect to make to the Pact Server
			dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String("/health"),
				Headers: dsl.MapMatcher{
					"Accept":       dsl.String("application/json"),
					"X-Request-Id": dsl.String(requestID),
				},
			},
		).
		WillRespondWith(
			// Configuring the response the Pact Server with response with
			dsl.Response{
				Status: http.StatusOK,
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.String("application/json"),
				},
				Body: dsl.Match(consumer.HealthResponse{}), // The example response values are configured in the production code using Tags
			},
		)

	testFunc := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		// The Pact standalone executables run a testFunc server to make testFunc requests against. It uses a random port to expose
		// a fake HTTP server to testFunc against base on the configured interactions.
		hc := consumer.NewHealthChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		hr := hc.Check(ctx) // Make an API call

		// Check that the client is decoding the returned JSON correctly
		assert.Equal(t, "OK", hr.Status)
		assert.Equal(t, int64(36), hr.Integer)
		assert.Equal(t, 12.34, hr.Float)
		assert.False(t, hr.Boolean)

		return nil
	}

	// This runs the testFunc
	assert.NoError(t, pact.Verify(testFunc))

	// Tears down the Pact server and produces the Pact file base on the inteactions the Pact Test Server received.
	//
	// NOTE: If you are running multiple tests again the same Provider run tear down at the end of the suite and not
	//	after each testFunc. Running teardown per testFunc will mean you get a pactfile per tests
	pact.Teardown()
}
