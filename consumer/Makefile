test-health:
	go test -v ./health

publish-health-pact:
	pact-broker publish pacts/health_checker_client-demo_app.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-i-deploy-health:
	pact-broker can-i-deploy --pacticipant 	"Health Checker Client" --broker-base-url http://broker:9393 --latest

test-thing:
	go test -v ./thing

publish-thing-pact:
	pact-broker publish pacts/thing_client-demo-app.json --consumer-app-version $$(uuid) --broker-base-url http://broker:9393

can-i-deploy-thing:
	pact-broker can-i-deploy --pacticipant "Thing Client" --broker-base-url http://broker:9393 --latest