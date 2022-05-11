##
## Build
##
FROM golang:1.18 AS build

WORKDIR /app
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=direct
RUN go mod download

COPY *.go ./

RUN go build -mod=vendor -o bin/server ./cmd/

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=build bin/server /server

EXPOSE 8000

CMD [ "/server" ]

