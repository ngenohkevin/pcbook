proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto
clean:
	rm -f pb/*.go
server:
	go run cmd/server/main.go
client:
	go run cmd/client/main.go
test:
	go test -cover -race ./...

.PHONY: proto clean run