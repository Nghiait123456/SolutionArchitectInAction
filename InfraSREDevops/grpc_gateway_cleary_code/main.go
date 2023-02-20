package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"grpc_gateway/gapi"
	"grpc_gateway/pb"
	"log"
	"net"
	"net/http"
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

func runGatewayServer() {
	server := gapi.NewServer()

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("start Http Gateway-grpc server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start Http Gateway-grpc server  server:", err)
	}
}

func main() {
	go runGatewayServer()
	runGrpcServer()
}
