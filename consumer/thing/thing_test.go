package thing_test

import (
	"github.com/pact-foundation/pact-go/dsl"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	pact := &dsl.Pact{
		Consumer: "Thing Client", // The name of the consumer using the API. In this cade out Health Checker Client
		Provider: "Demo App",     // The name of the Provider we are testing against
		LogLevel: "NONE",
	}

	status := m.Run()

	pact.Teardown()
	os.Exit(status)
}
