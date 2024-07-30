FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /go/src/accounts-service

COPY . .

FROM scratch

COPY --from=build /go/bin/accounts-service /go/bin/

EXPOSE 50014

ENV MIGRATE=false

ENTRYPOINT [ "/go/bin/accounts-service","-migrate=${MIGRATE}" ]