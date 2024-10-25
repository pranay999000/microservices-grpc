package main

import (
	"log"
	"net"

	"github.com/pranay999000/microservices-grpc/user-service/domain"
	pb "github.com/pranay999000/microservices-grpc/user-service/proto"
	"github.com/pranay999000/microservices-grpc/user-service/repository"
	"github.com/pranay999000/microservices-grpc/user-service/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userMap := make(map[int32]*domain.User)
	repo := repository.NewUserRepo(&userMap)

	grpcserver := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcserver, service.NewUserServiceServer(repo))

	log.Println("User service running on post :50051")
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}