package gapi

import "grpc_gateway/pb"

type Server struct {
	pb.UnimplementedSimpleBankServer
}

func NewServer() *Server {
	return &Server{}
}
