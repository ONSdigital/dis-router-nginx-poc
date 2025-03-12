FROM golang:1.24-bullseye AS base

ENV GOCACHE=/go/.go/cache GOPATH=/go/.go/path TZ=Europe/London

# Install github.com/cespare/reflex
RUN PATH=$PATH:/bin

# Clean cache, as we want all modules in the container to be under /go/.go/path
RUN go clean -modcache

COPY . /redirector

# Map between the working directories of dev and live
RUN ln -s /go /redirector
WORKDIR /redirector
