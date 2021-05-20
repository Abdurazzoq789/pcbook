gen:
	protoc -I=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 9090

client:
	go run cmd/client/main.go -address 0.0.0.0:9090

test:
	go test -cover -race ./...

.PHONY:	gen clean server client test