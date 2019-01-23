FROM golang:1.11-alpine3.8 AS build
WORKDIR /go/src/github.com/sabeersulaiman/hello-go
RUN apk add bash git gcc g++ libc-dev
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install

FROM alpine:3.8
WORKDIR /app
COPY --from=build /go/bin/hello-go hellogo
ENTRYPOINT [ "/app/hellogo" ]