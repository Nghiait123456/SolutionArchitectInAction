package gapi_balance_service

import (
	"grpc_basic_cleary_server/pb/pb_balance_service"
)

type Server struct {
	pb_balance_service.UnimplementedBalanceHandleServer
}

func NewServer() *Server {
	return &Server{}
}
