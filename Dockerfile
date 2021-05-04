FROM golang:latest

LABEL maintainer="Aviv"

WORKDIR /app

COPY go.mod .

COPY go.mod .

RUN go mod download

COPY . .

ENV port = 8080

RUN go build 

CMD ["./RestApi_Golang"]