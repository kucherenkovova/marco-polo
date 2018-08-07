proto:
	protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/*.proto

client: proto

server: proto

adapter: proto

all: client adapter server