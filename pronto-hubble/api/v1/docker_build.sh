#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build ./cmd/pronto-hubble-server

docker build -t pronto/hubble-server .