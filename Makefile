gen:
	protoc -I=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

cert:
	cd cert; bash gen.sh; cd ..

.PHONY:	gen clean server client test cert