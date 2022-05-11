package health_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"pact-consumer-demo/health"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

var pact *dsl.Pact

func setup() {
	// Instantiating Pact
	//
	// NOTE: If you are running multiple tests again the same Provider create one instance dls.Pact that is used by all tests.
	pact = &dsl.Pact{
		Consumer: "Health Checker Client", // The name of the consumer using the API. In this cade out Health Checker Client
		Provider: "Demo Health Endpoint",  // The name of the Provider we testing against
		LogLevel: "NONE",
		PactDir:  "../pacts",
	}
}

func TestHealthCheckWithStructPactExamples(t *testing.T) {
	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		// Providing the expectations for the Provider to setup
		Given("The service is up and running").
		// Describing the request that will be made
		UponReceiving("A GET request for the services health using Struct with Pact examples").
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
				// The example response values are configured in the production code using Tags
				Body: dsl.Match(health.ResponseWithPactExamples{}),
			},
		)

	testFunc := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		// The Pact standalone executables run a testFunc server to make testFunc requests against. It uses a random port to expose
		// a fake HTTP server to testFunc against base on the configured interactions.
		hc := health.NewChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		hr := hc.Check(ctx) // Make an API call

		// Check that the client is decoding the returned JSON correctly
		assert.Equal(t, "NOT OK", hr.Status)
		assert.Equal(t, int64(36), hr.Integer)
		assert.Equal(t, 12.34, hr.Float)
		assert.False(t, hr.Boolean)

		return nil
	}

	// This runs the testFunc
	assert.NoError(t, pact.Verify(testFunc))
}

func TestHealthCheckWithMapOfKeysAndMatcherValues(t *testing.T) {
	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		Given("The service is up and running").                                                      // Providing the expectations for the Provider to setup
		UponReceiving("A GET request for the services health using map of keys and matcher values"). // Describing the request that will be made
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
				Body: map[string]interface{}{
					"Status":  dsl.Like("OK"),
					"message": dsl.Like("Different message that the producer test"),
					"integer": dsl.Like(36),
					"float":   dsl.Like(12.34),
					"boolean": dsl.Like(true),
				}, // The test uses a hardcoded map of keys and matcher values
			},
		)

	testFunc := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		// The Pact standalone executables run a testFunc server to make testFunc requests against. It uses a random port to expose
		// a fake HTTP server to testFunc against base on the configured interactions.
		hc := health.NewChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		hr := hc.Check(ctx) // Make an API call

		// Check that the client is decoding the returned JSON correctly
		assert.Equal(t, "OK", hr.Status)
		assert.Equal(t, int64(36), hr.Integer)
		assert.Equal(t, 12.34, hr.Float)
		assert.True(t, hr.Boolean)

		return nil
	}

	// This runs the testFunc
	assert.NoError(t, pact.Verify(testFunc))
}

func TestHealthCheckWithMapOfKeysAndHardValues(t *testing.T) {
	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		Given("The service is up and running").                                       // Providing the expectations for the Provider to setup
		UponReceiving("A GET request for the services health using hard coded JSON"). // Describing the request that will be made
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
				Body: map[string]interface{}{
					"Status":  "OK",
					"message": "Testing Testing 123",
					"integer": 36,
					"float":   12.34,
					"boolean": true,
				}, // The test uses a hardcoded map of keys and values
			},
		)

	testFunc := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		// The Pact standalone executables run a testFunc server to make testFunc requests against. It uses a random port to expose
		// a fake HTTP server to testFunc against base on the configured interactions.
		hc := health.NewChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		hr := hc.Check(ctx) // Make an API call

		// Check that the client is decoding the returned JSON correctly
		assert.Equal(t, "OK", hr.Status)
		assert.Equal(t, int64(36), hr.Integer)
		assert.Equal(t, 12.34, hr.Float)
		assert.True(t, hr.Boolean)

		return nil
	}

	// This runs the testFunc
	assert.NoError(t, pact.Verify(testFunc))
}

func TestHealthCheckWithStructWithVariableSet(t *testing.T) {
	const requestID = "123456789-qwerty"

	pact.AddInteraction().
		Given("The service is up and running").                                                // Providing the expectations for the Provider to setup
		UponReceiving("A GET request for the services health using struct with variable set"). // Describing the request that will be made
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
				Body: dsl.Like(health.ResponseWithoutPactExamples{
					Status:  "OK",
					Integer: 36,
					Float:   12.34,
					Boolean: false,
				}), // The test uses a hardcoded string of JSON to create the response
			},
		)

	testFunc := func() error {
		ctx := context.WithValue(context.Background(), "request-id", requestID)

		// The Pact standalone executables run a testFunc server to make testFunc requests against. It uses a random port to expose
		// a fake HTTP server to testFunc against base on the configured interactions.
		hc := health.NewChecker(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
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
}

func shutdown() {
	// Tears down the Pact server and produces the Pact file base on the inteactions the Pact Test Server received.
	//
	// NOTE: If you are running multiple tests again the same Provider run tear down at the end of the suite and not
	//	after each testFunc. Running teardown per testFunc will mean you get a pactfile per tests
	pact.Teardown()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
