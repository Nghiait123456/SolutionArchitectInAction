package gapi_auth_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_basic_cleary_server/pb/pb_auth_service"
	"time"
)

func (server *Server) CreateUser(ctx context.Context, req *pb_auth_service.CreateUserRequest) (*pb_auth_service.CreateUserResponse, error) {
	if req.GetUsername() == "admin" {
		return &pb_auth_service.CreateUserResponse{
			User: &pb_auth_service.User{
				Username:          "admin",
				FullName:          "admin",
				Email:             "admin@admin",
				PasswordChangedAt: nil,
				CreatedAt:         nil,
			},
		}, status.Errorf(codes.OK, "done")
	}

	return nil, status.Errorf(codes.Internal, "internal errors")
}

func (server *Server) CreateUserSample(ctx context.Context, req *pb_auth_service.CreateUserRequest) (*pb_auth_service.CreateUserResponse, error) {
	// bind DTO
	// validate
	// call core handle
	// convert status grpc

	if req.GetUsername() == "admin" {
		return &pb_auth_service.CreateUserResponse{
			User: &pb_auth_service.User{
				Username:          "admin",
				FullName:          "admin",
				Email:             "admin@admin",
				PasswordChangedAt: timestamppb.New(time.Unix(123, 0)),
				CreatedAt:         timestamppb.New(time.Unix(456, 0)),
			},
		}, status.Errorf(codes.OK, "done")
	}

	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
