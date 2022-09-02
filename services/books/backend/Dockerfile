# syntax=docker/dockerfile:1

FROM golang:1.19.0-alpine3.16

WORKDIR /books

COPY go.sum ./
COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /pl-books-v1

CMD [ "/pl-books-v1" ]