package gapi

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_gateway/pb"
	"time"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("have new request, req= ", req)
	if req.GetUsername() == "admin" {
		return &pb.CreateUserResponse{
			User: &pb.User{
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

func (server *Server) CreateUserSample(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// bind DTO
	// validate
	// call core handle
	// convert status grpc

	if req.GetUsername() == "admin" {
		return &pb.CreateUserResponse{
			User: &pb.User{
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
