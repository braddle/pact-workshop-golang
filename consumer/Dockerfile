FROM golang:latest as build

WORKDIR /code

RUN curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v1.88.81/pact-1.88.81-linux-x86_64.tar.gz
RUN tar xzf pact-1.88.81-linux-x86_64.tar.gz -C /
ENV PATH $PATH:/pact/bin

CMD tail -f /dev/null

