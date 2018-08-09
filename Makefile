SHELL := /bin/bash

all: protoc client server adapter

protoc: proto/services.proto proto/messages.proto
	protoc -I ./proto --go_out=plugins=grpc:./proto proto/*.proto

client: services/client/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/client services/client/main.go

server: services/server/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/server services/server/main.go

adapter: services/adapter/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/adapter services/adapter/main.go
