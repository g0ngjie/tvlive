#!/bin/sh

# 版本号
VERSION=0.0.1

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./__run-amd64-${VERSION}.exe main.go

# MacOS 64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./__run-darwin-amd64-${VERSION} main.go
# arm
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./__run-darwin-arm64-${VERSION} main.go

# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./__run-linux-amd64-${VERSION} main.go
# arm64 64位架构 ARMv8 适用于ARMv8设备
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./__run-linux-arm64-${VERSION} main.go
# arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./__run-linux-arm-${VERSION} main.go