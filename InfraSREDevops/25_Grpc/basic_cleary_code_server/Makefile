
proto-generate:
	rm -rf pb/*.go
	rm -rf pb/auth_service/*.go
	protoc --proto_path=proto/auth_service --go_out=pb/pb_auth_service --go_opt=paths=source_relative \
        --go-grpc_out=pb/pb_auth_service --go-grpc_opt=paths=source_relative \
        proto/auth_service/*.proto

server:
	go run main.go

make-evans:
	evans --host localhost --port 9090 -r repl