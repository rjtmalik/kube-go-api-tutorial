# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
RUN apk add build-base
# The line RUN apk add build-base was required only for the sqlite. It can be removed on a real db

WORKDIR  /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# to maintain the directory structure, you need to mention folder name in the destination also
COPY db/ ./db/
COPY handlers/ ./handlers/
COPY repository/ ./repository/
COPY repository_interfaces/ ./repository_interfaces/
COPY *.go ./
RUN go build -o /api_blank

EXPOSE 8001

CMD ["/api_blank"]