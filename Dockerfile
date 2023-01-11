FROM golang:1.19 as setup

ENV GO111MODULE=on

RUN apt-get update && \
    apt-get upgrade -y

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

FROM setup as build
CMD ["go", "build", "."]

FROM alpine:latest as runner
EXPOSE 8080
COPY --from=build pbl-back /pbl-back
CMD ["/pbl-back"]
