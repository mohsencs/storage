# syntax=docker/dockerfile:1

FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

#ENTRYPOINT ["./main"]
FROM gcr.io/distroless/base-debian11

WORKDIR /src

COPY --from=build ./app .

#USER root:root

ENTRYPOINT ["./main"]