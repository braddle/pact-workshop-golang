test-health:
	go test ./health_test.go

publish-health-pact:
	pact-broker publish pacts/healthchecker-demohealth.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393
