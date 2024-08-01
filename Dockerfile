FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /go/src/accounts-service

COPY . .

RUN go mod download

ARG CGO_ENABLED=0 GOOS=linux

RUN go build -o /out/accounts-service

FROM alpine:3.20

WORKDIR /app

COPY --from=build /out/accounts-service ./

EXPOSE 50014

ENTRYPOINT [ "/app/accounts-service", "-migrate", "$MIGRATE" ]