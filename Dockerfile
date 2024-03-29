## Build
FROM golang:1.21-alpine AS buildenv

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

ADD . .

ENV GO111MODULE=on

RUN  go build -o main cmd/main.go

## Deploy
FROM alpine

WORKDIR /

COPY --from=buildenv  /app/ /

EXPOSE 9191

CMD ["/main"]