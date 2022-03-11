# Consumer Driven Contract Testing with Pact Lab (Golang)

## Introduction 

The purpose of this Lab exercise it to introduce you to Consuner Driven Contract testing with [Pact](https://pact.io/).

### What we will cover

Duing this lab we will look at Using Pact to create a contract between a HTTP REST API and a client that uses the API. 
The Lab is broken down into 3 parts:

1. We will look at concepts and language introduced by Pact. 
2. We will look at a small example project to explore how to use Pact with Golang. 
3. You will create a contract for a single GET API endpoint.

### Things we will not be covering

- Using Data Providers to populate a datasource to enable dynamic testing of the Provider
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

This will start 4 seperate contains:

1. Running a small HTTP REST API with two endpoints `/health` & `/thing/{id}`. (The Provider)
2. A Go environment with an example API client to interact with the `/health` endpoint
3. A Pact Broker 
4. Database for the Pact Broker (PostgreSQL) 

### Exploring the Example Project

#### The Consumer

The Consumer is the code that interact with an API. I have already created an [example Consumer](consumer/health.go) for 
the `/health` endpoint of our Provider. The [code](consumer/health.go) is commented to explain what is happening at all 
key point of execution, spend a few minutes exploring it to familiarise yourself with it before continuing. 

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

The testing of our Consumer is now complete but we have not shared the Pact file (The Contract) with the Producer. 
This is where we use the Pact Broker.

## Pact Broker

[pact broker](http://localhost:9393/)

## Producer

`make jump-to-consumer`
