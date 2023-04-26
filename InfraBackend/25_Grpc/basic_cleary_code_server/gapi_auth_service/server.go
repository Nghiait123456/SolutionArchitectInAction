package gapi_auth_service

import "grpc_basic_cleary_server/pb/pb_auth_service"

type Server struct {
	pb_auth_service.UnimplementedSimpleBankServer
}

func NewServer() *Server {
	return &Server{}
}
