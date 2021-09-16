FROM golang:1.16-alpine

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/boazsoftwares/Boaz.Api.Go/cmd/boazApi
RUN apk update && apk add git
RUN cd /build && git clone https://github.com/boazsoftwares/Boaz.Api.Go.git

RUN cd /build/Boaz.Api.Go/cmd/boazApi && go build

EXPOSE 8080

ENTRYPOINT [ "Boaz.Api.Go/cmd/boazApi/main" ]