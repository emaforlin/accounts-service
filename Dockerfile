FROM golang:1.22.5-alpine3.20 AS build-stage

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /accounts-service


FROM alpine:3.20 AS build-release-stage

WORKDIR /

COPY --from=build-stage /accounts-service /accounts-service

COPY  entrypoint.sh entrypoint.sh

EXPOSE 50014

ENV MIGRATE=false

RUN chmod +x entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]