package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_basic_cleary_server/gapi_auth_service"
	"grpc_basic_cleary_server/gapi_balance_service"
	"grpc_basic_cleary_server/pb/pb_auth_service"
	"grpc_basic_cleary_server/pb/pb_balance_service"
	"log"
	"net"
)

func runGrpcServer() {
	serverAuth := gapi_auth_service.NewServer()
	serverBalance := gapi_balance_service.NewServer()
	grpcServer := grpc.NewServer()

	pb_auth_service.RegisterSimpleBankServer(grpcServer, serverAuth)
	pb_balance_service.RegisterBalanceHandleServer(grpcServer, serverBalance)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func main() {
	runGrpcServer()
}
