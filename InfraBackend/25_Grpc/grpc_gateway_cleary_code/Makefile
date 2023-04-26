
proto-generate:
	rm -rf pb/*.go
	protoc --proto_path=proto/auth --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative\
         proto/auth/*.proto

server:
	go run main.go

make-evans:
	evans --host localhost --port 9090 -r repl