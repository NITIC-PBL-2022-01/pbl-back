FROM golang:1.19 as setup

RUN apt-get update && \
    apt-get upgrade -y

ADD . /src
WORKDIR /src

FROM setup as build
CMD ["go", "build", "."]

FROM alpine:latest as runner
EXPOSE 8080
COPY --from=build pbl-back /pbl-back
CMD ["/pbl-back"]
