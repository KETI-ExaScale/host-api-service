#!/bin/bash

dockerID="ketidevit2"

go mod tidy
go mod vendor


go build -o build/bin/host-api-service cmd/main.go
rm -rf /usr/lib/systemd/system/host-api-service.service
rm -rf /usr/bin/local/.keti
cp --force build/host-api-service.service /usr/lib/systemd/system/host-api-service.service
mkdir /usr/bin/local/.keti
cp --force build/bin/host-api-service /usr/bin/local/.keti/host-api-service

systemctl daemon-reload
systemctl restart host-api-service
systemctl enable host-api-service