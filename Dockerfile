FROM golang:1.15.2 AS build
ADD . /src
WORKDIR /src
ENV GO111MODULE=on

RUN make dep

ENTRYPOINT ["/src/entrypoint.sh"]
