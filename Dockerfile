FROM golang:1.22.6-alpine3.20

RUN mkdir app

COPY . /app

WORKDIR /app

EXPOSE 8080

