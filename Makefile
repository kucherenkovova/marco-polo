SHELL := /bin/bash
PROJECT=/go/src/github.com/kucherenkovova/marco-polo.git

all: protoc client server adapter

protoc: proto/services.proto proto/messages.proto
	protoc -I ./proto --go_out=plugins=grpc:./proto proto/*.proto

client: services/client/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/client services/client/main.go

server: services/server/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/server services/server/main.go

adapter: services/adapter/main.go proto/services.pb.go proto/messages.pb.go
	go build -o bin/adapter services/adapter/main.go

images:
	docker build --build-arg "SERVICE=adapter" --build-arg "PROJECT=$(PROJECT)" -t mikefaraponov/marco-polo-adapter .
	docker build --build-arg "SERVICE=client" --build-arg "PROJECT=$(PROJECT)" -t mikefaraponov/marco-polo-client .
	docker build --build-arg SERVICE=server --build-arg "PROJECT=$(PROJECT)" -t mikefaraponov/marco-polo-server .