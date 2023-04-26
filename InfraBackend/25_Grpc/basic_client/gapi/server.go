package gapi

import "grpc_basic/pb"

type Server struct {
	pb.UnimplementedSimpleBankServer
}

func NewServer() *Server {
	return &Server{}
}
