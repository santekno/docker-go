FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/docker-go main.go

ENTRYPOINT ["/app/docker-go"]