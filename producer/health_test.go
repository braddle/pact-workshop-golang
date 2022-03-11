package main

import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealth(t *testing.T) {
	pact := dsl.Pact{
		Provider: "DemoHealth",
		LogLevel: "NONE",
	}

	_, err := pact.VerifyProvider(
		t,
		types.VerifyRequest{
			ProviderBaseURL:            "http://localhost:8082",
			BrokerURL:                  "http://broker:9393",
			Provider:                   "DemoHealth",
			FailIfNoPactsFound:         true,
			PublishVerificationResults: true,
			ProviderVersion:            "1.0.0",
		},
	)

	assert.NoError(t, err)
}
