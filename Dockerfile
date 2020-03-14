ARG GO_VERSION=1.14.0
FROM golang:${GO_VERSION}-alpine

WORKDIR /app

RUN apk add bash ca-certificates git openssh build-base curl

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ARG APP_NAME
RUN go build -o main ${APP_NAME}/*

CMD ["./main"]
