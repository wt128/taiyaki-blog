FROM golang:1.20.0-alpine

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./main.go ${ROOT}
RUN go mod init
COPY ./go.mod ${ROOT}

RUN go mod tidy
EXPOSE 8080