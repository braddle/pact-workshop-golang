package main

import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealth(t *testing.T) {
	// Instantiating Pact
	pact := dsl.Pact{
		Provider: "DemoHealth",
		LogLevel: "NONE",
	}

	// Verify we meet the contract in the Pact file
	_, err := pact.VerifyProvider(
		t,
		types.VerifyRequest{
			ProviderBaseURL:            "http://localhost:8082", // The URL the service under test is running on
			BrokerURL:                  "http://broker:9393",    // The URL of the Pact Broker to obtain Pact Files from
			Provider:                   "DemoHealth",            // The name of this service. This need to be the match the one set when the provider publishes the Pact File
			FailIfNoPactsFound:         true,
			PublishVerificationResults: true, // Flag the mean result of test run are published to the broker. Usually on set to true in CI/CD pipeline
			ProviderVersion:            "1.0.0",
		},
	)

	assert.NoError(t, err)
}
