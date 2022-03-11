test-health:
	go test -v ./health_test.go

publish-health-pact:
	pact-broker publish pacts/health_checker_client-demo_health_endpoint.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-deploy-health:
	pact-broker can-i-deploy --pacticipant 	"Health Checker Client" --broker-base-url http://broker:9393 --latest

