FROM golang:1.23-alpine

WORKDIR /app/src

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src .
