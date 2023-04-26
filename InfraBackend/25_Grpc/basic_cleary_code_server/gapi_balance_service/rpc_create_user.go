package gapi_balance_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_basic_cleary_server/pb/pb_balance_service"
)

func (server *Server) RequestBalance(ctx context.Context, rq *pb_balance_service.BalanceRequest) (*pb_balance_service.BalanceResponse, error) {
	if rq.GetFullName() == "admin" {
		return &pb_balance_service.BalanceResponse{
			Balance: &pb_balance_service.Balance{
				Amount: 10000,
			},
		}, status.Errorf(codes.OK, "done")
	}

	return nil, status.Errorf(codes.Internal, "internal errors")
}
