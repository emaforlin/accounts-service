FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /go/src/accounts-service

COPY . .

RUN go build -o /go/bin/accounts-service .

FROM scratch

COPY --from=build /go/bin/accounts-service /go/bin/accounts-service

EXPOSE 50014

ENV MIGRATE=false

ENTRYPOINT [ "/go/bin/accounts-service","-migrate=${MIGRATE}" ]