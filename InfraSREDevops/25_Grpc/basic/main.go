package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_basic/gapi"
	"grpc_basic/pb"
	"log"
	"net"
)

func runGrpcServer() {
	server := gapi.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
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
