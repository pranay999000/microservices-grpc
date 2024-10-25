package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/pranay999000/microservices-grpc/post-service/domain"
)

type PostRepo struct {
	PostsMap	*map[int32]*domain.Post
}

func NewPostRepo(postMap *map[int32]*domain.Post) *PostRepo {
	return &PostRepo{
		PostsMap: postMap,
	}
}

func (repo *PostRepo) CreatePost(ctx context.Context, post *domain.Post) error {

	var err error

	if _, ok := (*repo.PostsMap)[post.Id]; ok {
		err = errors.New("id already exists")
		log.Println("failed to create post", err)
		return err
	}

	(*repo.PostsMap)[post.Id] = post
	return nil

}

func (repo *PostRepo) GetPostById(ctx context.Context, postId int32) (*domain.Post, error) {

	var err error

	if _, ok := (*repo.PostsMap)[postId]; !ok {
		err = errors.New("post with id " + strconv.Itoa(int(postId)) + " does not exist")
		log.Println("failed to get post", err)
		return nil, err
	}

	return (*repo.PostsMap)[postId], nil

}

func (repo *PostRepo) GetAllPosts(ctx context.Context) []*domain.Post {
	resp := []*domain.Post{}

	for _, value := range (*repo.PostsMap) {
		resp = append(resp, value)
	}

	return resp
}