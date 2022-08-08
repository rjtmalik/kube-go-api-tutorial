# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR  /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /api_blank

EXPOSE 8000

CMD ["/api_blank"]