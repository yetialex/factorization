FROM golang:1.13-alpine

COPY . /src/app

WORKDIR /src/app

RUN env GOOS=linux GOARCH=amd64 go build -o ./dist/prime_factorization ./cmd/prime_factorization
