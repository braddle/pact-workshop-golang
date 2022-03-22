package consumer_test

import (
	"github.com/pact-foundation/pact-go/dsl"
	"os"
	"testing"
)

var pact *dsl.Pact

func TestMain(m *testing.M) {
	pact = &dsl.Pact{
		Consumer: "Health Checker Client", // The name of the consumer using the API. In this cade out Health Checker Client
		Provider: "Demo Health Endpoint",  // The name of the Provider we testing against
		LogLevel: "NONE",
	}

	status := m.Run()

	pact.Teardown()
	os.Exit(status)
}
