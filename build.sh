#!/usr/bin/env bash
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
swag init cmd/handlers/main.go -o cmd/handlers/docs -g cmd/main.go
go mod tidy

MODULE_NAME=$(grep ^module go.mod | cut -d " " -f2)
go build -o bin/api -ldflags="-s -w -X ${MODULE_NAME}/internal/config/global.CommitHash=$(git rev-parse HEAD)" ./cmd