# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./


RUN go build -o /weight_logger_backend
EXPOSE 10000
CMD [ "/weight_logger_backend" ]
