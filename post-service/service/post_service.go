package service

import (
	"context"

	"github.com/pranay999000/microservices-grpc/post-service/domain"
	pb "github.com/pranay999000/microservices-grpc/post-service/proto"
	"github.com/pranay999000/microservices-grpc/post-service/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostServiceServer struct {
	repo *repository.PostRepo
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(repo *repository.PostRepo) *PostServiceServer {
	return &PostServiceServer{
		repo: repo,
	}
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	
	if req.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "bad request %d", req.Id)
	}

	post, err := s.repo.GetPostById(ctx, req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &pb.PostResponse{
		Id: post.Id,
		Title: post.Title,
		Content: post.Content,
	}, nil

}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {

	if req.Id < 1 || req.Title == "" || req.Content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bad request: %d, %s, %s", req.Id, req.Title, req.Content)
	}

	err := s.repo.CreatePost(ctx, &domain.Post{
		Id: req.Id,
		Title: req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}

	return &pb.CreatePostResponse{
		Message: "successfully create post",
	}, nil

}

func (s *PostServiceServer) Posts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {

	posts := s.repo.GetAllPosts(ctx)

	var postResp []*pb.PostResponse

	for _, post := range posts {
		postResp = append(postResp, &pb.PostResponse{
			Id: post.Id,
			Title: post.Title,
			Content: post.Content,
		})
	}

	return &pb.GetAllPostsResponse{
		Posts: postResp,
	}, nil

}