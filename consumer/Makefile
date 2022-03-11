test-health:
	go test -v ./health_test.go

publish-health-pact:
	pact-broker publish pacts/healthchecker-demohealth.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-deploy-health:
	pact-broker can-i-deploy --pacticipant HealthChecker --broker-base-url http://broker:9393 --latest

