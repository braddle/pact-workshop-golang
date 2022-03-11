# Consumer Driven Contract Testing with Pact Lab (Golang)

## Introduction 

The purpose of this Lab exercise it to introduce you to Consumer Driven Contract testing with [Pact](https://pact.io/).

### What we will cover

During this lab we will look at Using Pact to create a contract between a HTTP REST API and a client that uses the API. 
The Lab is broken down into 3 parts:

1. We will look at concepts and language introduced by Pact. 
2. We will look at a small example project to explore how to use Pact with Golang. 
3. You will create a contract for a single GET API endpoint.

### Things we will not be covering

- Using [Provider State](https://docs.pact.io/getting_started/terminology#provider-state) to populate a datasource to 
  enable dynamic testing of the Provider
- Pact Contracts for Async communication
- Setting up Pact in a CI/CD Pipeline

### Prerequisite

- Docker
- Docker-Compose
- Knowledge of how to create and interact with HTTP REST APIs in Golang
  
## Our Project

The code we will look at, and you will create is setup to run in Docker containers to allow you to concentrate on looking
a Pact and not having to install and configure you own machine to run Pact.

### Getting Started

To start the docker containers using the following Make target

```shell
make start
```

This will start 4 separate containers:

1. Running a small HTTP REST API with two endpoints `/health` & `/thing/{id}`. (The Provider)
2. A Go environment with an example API client to interact with the `/health` endpoint
3. A Pact Broker 
4. Database for the Pact Broker (PostgreSQL) 

### Exploring the Example Project

#### The Consumer

The [Consumer](https://docs.pact.io/getting_started/terminology#service-consumer) is the code that interacts with an API. 

I have already created an [example Consumer](consumer/health.go) for the `/health` endpoint of our Provider. 
The [code](consumer/health.go) is commented to explain what is happening at all key point of execution, spend a few 
minutes exploring it to familiarise yourself with it before continuing. 

The code assumes that Making the HTTP call and decoding the JSON response is successful because you would not use Pact 
to test these bits of the code. Handling these errors should be covered by separate unit tests.

The [tests](consumer/health_test.go) that I used to drive out the implementation cover the expected response from the 
`/health` endpoint when the service is running and healthy. The tests are commented to explain how to use Pacts 
[Golang Test Framework](https://github.com/pact-foundation/pact-go). Why not read through the tests now before we start 
to run them.

To run the tests you need to be on the consumer docker container. You can get access to Bash on the container using the 
following Make target

```shell
make jump-to-consumer
```

To run the tests for the [Health Check Client](consumer/health.go) you can use the following command 

```shell
make test-health 
```

This will run the tests against the Health Check Client and if they are successful a 
[Pact File](consumer/pacts/health_checker_client-demo_health_endpoint.json) will be created. It should look something like this:

```json
{
  "consumer": {
    "name": "Health Checker Client"
  },
  "provider": {
    "name": "Demo Health Endpoint"
  },
  "interactions": [
    {
      "description": "A GET request for the services health",
      "providerState": "The service is up and running",
      "request": {
        "method": "GET",
        "path": "/health",
        "headers": {
          "Accept": "application/json",
          "X-Request-Id": "123456789-qwerty"
        }
      },
      "response": {
        "status": 200,
        "headers": {
          "Content-Type": "application/json"
        },
        "body": {
          "Status": "OK",
          "boolean": false,
          "float": 12.34,
          "integer": 36
        },
        "matchingRules": {
          "$.body.Status": {
            "match": "type"
          },
          "$.body.boolean": {
            "match": "type"
          },
          "$.body.float": {
            "match": "type"
          },
          "$.body.integer": {
            "match": "type"
          }
        }
      }
    }
  ],
  "metadata": {
    "pactSpecification": {
      "version": "2.0.0"
    }
  }
}
```

##### Step by Step Video
![Tutorial video on writing consumer pact tests](https://www.youtube.com/watch?v=SCndSvUBlnw)

The testing of our Consumer is now complete but we have not shared the Pact file (The Contract) with the Producer. 
This is where we use the Pact Broker.

#### Pact Broker

The [Pact Broker](https://docs.pact.io/getting_started/sharing_pacts) is the repository for sharing Pact files between Consumers and Producers.

The home page of the [Pact Broker](http://localhost:9393/) shows a list of all Pacts that have been register and their status.

When you first load the [Pact Broker](http://localhost:9393/) you should see an Example Pact. Lets upload the Pact file 
we created for the Health Checker to the Broker. To does this you can use to command below on the Consumer Docker 
container.

```shell
make publish-health-pact
```

This will push the pact file to the broker, and it should not be visible on the [Pact Broker home page](http://localhost:9393/).

![Screen shot of the pact broker homepage showing the example and healthchecker pacts](docs/broker.png)

Before moving on explore the Pact Broker click through to the [Consumer](http://localhost:9393/pacticipants/HealthChecker) 
or the [Provider](http://localhost:9393/pacticipants/DemoHealth) look at the [Network Graph](http://localhost:9393/pacticipants/HealthChecker/network)

![Screen shot of the pact broker network graph between the Health Client and the Health endpoint](docs/network-graph.png)

![Video tutorial of sending Pact file to a Pact Broker](https://youtu.be/y-jW8dInFc4)

For the Lab we are running our own Pact Broker in Docker. 
If you do not want to administer you own Pact Broker you can use [Pactflow](https://pactflow.io/).

##### Can I Deploy

Now that we have some Consumer tests that produce a contract we need to know that our Producer meets the contract before
we are safe to deploy our Consumer that can be done with the `Can I Deploy` feature of the
[Pact standalone executables](https://github.com/pact-foundation/pact-ruby-standalone)

Before we do anything with the Producer let see if the we should deploy the Consumer. Try running the following command 
on the Consumer Docker container

```shell
make can-deploy-health
```

This should return a non-zero exit code saying the Pact has not been verified

![Screen shot of a failure response from can-i-deploy](docs/failed-can-i-deploy.png)

#### Producer

The [Producer](https://docs.pact.io/getting_started/terminology#service-provider) in the HTTP API providing the endpoints 
defined in the Pact File. 

Now that we have push our Pact File to the Pact Broker we can test the contract against the implementation. To run the 
test you first need to be on the Producer Docker container, you can do that by running the command below:

```shell
make jump-to-producer
```

We can now look at our [Producer Tests](producer/health_test.go). Take some time to look over the test code, the 
Provider tests a pretty small there is not much to setup because most information for them to run is in the Pact file.

To run the Producer tests for the `/health` endpoint run the follwing command on the Producer Docker container

```shell
make test-health
```

Once the tests have run go back to the [Pact Broker](http://localhost:9393) and see how the state of the Pact has been 
updated.

![Screenshot of the verified Pact on the Broker](docs/broker-verified.png)

##### Can I Deploy

Now that we have verified the Pact file from the Producer let see if the `Can I Deploy` tool thinks we can deploy the 
Consumer code. Try running the following command on the Consumer Docker container

```shell
make can-deploy-health
```

This now returns a zero exit code saying that the Pact is verified

![Screen shot of a success response from can-i-deploy](docs/successful-can-i-deploy.png)




