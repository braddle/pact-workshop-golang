test-health:
	go test -v ./health

publish-health-pact:
	pact-broker publish pacts/health_checker_client-demo_health_endpoint.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-i-deploy-health:
	pact-broker can-i-deploy --pacticipant 	"Health Checker Client" --broker-base-url http://broker:9393 --latest

test-thing:
	go test -v ./thing_test.go

publish-thing-pact:
	pact-broker publish pacts/thing_client-thing_endpoint.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-i-deploy-thing:
	pact-broker can-i-deploy --pacticipant 	"Thing Client" --broker-base-url http://broker:9393 --latest