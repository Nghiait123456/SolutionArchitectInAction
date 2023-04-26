package gapi

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_gateway/pb"
	"time"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if req.GetUsername() == "admin" {
		return &pb.LoginUserResponse{
			User: &pb.User{
				Username:          "admin",
				FullName:          "admin",
				Email:             "admin",
				PasswordChangedAt: nil,
				CreatedAt:         nil,
			},
			SessionId:             "1234",
			AccessToken:           "12345",
			RefreshToken:          "12345",
			AccessTokenExpiresAt:  timestamppb.New(time.Unix(123, 0)),
			RefreshTokenExpiresAt: timestamppb.New(time.Unix(456, 0)),
		}, nil
	}
	return nil, status.Errorf(codes.Internal, "internal error")
}
