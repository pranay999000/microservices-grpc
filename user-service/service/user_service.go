package service

import (
	"context"

	"github.com/pranay999000/microservices-grpc/user-service/domain"
	pb "github.com/pranay999000/microservices-grpc/user-service/proto"
	"github.com/pranay999000/microservices-grpc/user-service/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	repo 	*repository.UserRepo
	pb.UnsafeUserServiceServer
}

func NewUserServiceServer(repo *repository.UserRepo) *UserServiceServer {
	return &UserServiceServer{
		repo: repo,
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	
	if req.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "bad request %d", req.Id)
	}

	user, err := s.repo.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}

	return &pb.UserResponse{
		Id: user.Id,
		Name: user.Name,
	}, nil

}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	if req.Id < 1 || req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bad request: %d, %s", req.Id, req.Name)
	}

	err := s.repo.CreateUser(ctx, &domain.User{
		Id: req.Id,
		Name: req.Name,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &pb.CreateUserResponse{
		Message: "successfully created user",
	}, nil

}
