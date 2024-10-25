package main

import (
	"log"
	"net"

	"github.com/pranay999000/microservices-grpc/post-service/domain"
	pb "github.com/pranay999000/microservices-grpc/post-service/proto"
	"github.com/pranay999000/microservices-grpc/post-service/repository"
	"github.com/pranay999000/microservices-grpc/post-service/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	postMap := make(map[int32]*domain.Post)
	repo := repository.NewPostRepo(&postMap)

	grpcserver := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcserver, service.NewPostServiceServer(repo))

	log.Println("Post service running on post :50052")
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}