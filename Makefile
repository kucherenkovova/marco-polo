all: protoc client server adapter

protoc: proto/services.proto proto/messages.proto
	protoc -I ./proto --go_out=plugins=grpc:./proto proto/*.proto

client: services/client/main.go proto/services.pb.go proto/messages.pb.go
	go build -o client services/client/main.go

server: services/server/main.go proto/services.pb.go proto/messages.pb.go
	go build -o server services/server/main.go

adapter: services/adapter/main.go proto/services.pb.go proto/messages.pb.go
	go build -o adapter services/adapter/main.go
