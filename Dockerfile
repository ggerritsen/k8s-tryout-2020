ARG GO_VERSION=1.14.0

FROM golang:${GO_VERSION}-alpine

ARG APP_NAME

RUN apk add --no-cache bash ca-certificates git openssh build-base curl

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ${APP_NAME}/*

CMD ["./main"]
