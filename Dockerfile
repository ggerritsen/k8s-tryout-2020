ARG GO_VERSION=1.14.0
FROM golang:${GO_VERSION}-alpine

WORKDIR app

RUN apk add bash ca-certificates git openssh build-base curl && \
    GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
