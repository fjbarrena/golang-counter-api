# syntax=docker/dockerfile:1

FROM golang:1.20.10-alpine3.17 AS build

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .
COPY controllers ./controllers
COPY models ./models 
COPY services ./services

RUN go get counter-api
RUN go build -o /counter-api

FROM scratch AS deploy

WORKDIR /

COPY --from=build /counter-api /counter-api

ENV GOPORT 40000
ENV GIN_MODE release

ENTRYPOINT ["/counter-api"]