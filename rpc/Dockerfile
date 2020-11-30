# Dockerfile.protogen
FROM golang:latest

LABEL maintainer="Ben Grewell <ben@grewelltech.com>"

ENV PROTOC_VERSION 3.6.1
ENV PROTOC_GEN_GO_VERSION v1.2.0

WORKDIR /go/src/github.com/BGrewell/perspective/rpc/

RUN apt update
RUN apt install -y protobuf-compiler python3 python3-pip
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN pip3 install grpcio-tools
RUN export PATH=$PATH:$GOPATH/bin
RUN echo $PATH

COPY proto proto

RUN mkdir -p grpc/go
RUN protoc -I proto/ --go_out=plugins=grpc:grpc/go proto/perspective.proto

RUN mkdir -p grpc/python
RUN python3 -m grpc_tools.protoc -I proto/ --python_out=grpc/python/. --grpc_python_out=grpc/python/. proto/perspective.proto