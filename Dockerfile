# syntax=docker/dockerfile:1
FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /gobochiotte Gobochiotte/cmd

## Deploy
FROM scratch
COPY --from=build /gobochiotte /gobochiotte
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/gobochiotte"]