FROM golang:1.17-alpine

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod .
COPY main.go .

RUN GOOS=linux GOARCH=amd64 go build -o action main.go

ENTRYPOINT ["/argo-appset-gitgen"]
