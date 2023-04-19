# syntax=docker/dockerfile:1

## Build step
FROM golang:1.20-alpine AS build

WORKDIR /app

ENV CGO_ENABLED 0

COPY go.* ./
COPY *.go ./

RUN go build -o /nsc-go-app

## Deploy step
FROM gcr.io/distroless/base-debian10 as prod

WORKDIR /

COPY --from=build /nsc-go-app /nsc-go-app

EXPOSE 80

ENTRYPOINT ["/nsc-go-app"]