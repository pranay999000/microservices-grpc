package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pranay999000/microservices-grpc/api-gateway/domain"
	postpb "github.com/pranay999000/microservices-grpc/post-service/proto"
	userpb "github.com/pranay999000/microservices-grpc/user-service/proto"
	"google.golang.org/grpc"
)

var (
	userClient userpb.UserServiceClient
	postClient postpb.PostServiceClient
)

func getUser(w http.ResponseWriter, r *http.Request) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	
	defer cancel()

	reqest_id := r.URL.Query().Get("id")
	user_id, err := strconv.Atoi(reqest_id)

	if err != nil {
		http.Error(w, "Provide a valid user id", http.StatusBadRequest)
		return
	}

	res, err := userClient.GetUser(ctx, &userpb.UserRequest{Id: int32(user_id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func createUser(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.UserRequest); err != nil {
		http.Error(w, "unable to decode user request", http.StatusBadRequest)
		return
	}

	res, err := userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Id: domain.UserRequest.Id,
		Name: domain.UserRequest.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func getPostById(w http.ResponseWriter, r *http.Request) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reqest_id := r.URL.Query().Get("id")
	post_id, err := strconv.Atoi(reqest_id)

	if err != nil {
		http.Error(w, "Provide a valid post id", http.StatusBadRequest)
		return
	}

	res, err := postClient.GetPost(ctx, &postpb.PostRequest{Id: int32(post_id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func getAllPosts(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := postClient.Posts(ctx, &postpb.GetAllPostsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func createPost(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&domain.PostRequest); err != nil {
		http.Error(w, "Unable to decode json file", http.StatusInternalServerError)
		return
	}

	res, err := postClient.CreatePost(ctx, &postpb.CreatePostRequest{
		Id: domain.PostRequest.Id,
		Title: domain.PostRequest.Title,
		Content: domain.PostRequest.Content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func main() {

	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
		return
	}

	defer userConn.Close()

	userClient = userpb.NewUserServiceClient(userConn)

	postConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to post service: %v", err)
		return
	}

	defer postConn.Close()

	postClient = postpb.NewPostServiceClient(postConn)
	userClient = userpb.NewUserServiceClient(userConn)

	http.HandleFunc("/user", getUser)
	http.HandleFunc("/user/create", createUser)
	
	http.HandleFunc("/post", getPostById)
	http.HandleFunc("/posts", getAllPosts)
	http.HandleFunc("/post/create", createPost)


	log.Printf("API Gatway running on port: %d", 3000)
	log.Fatal(http.ListenAndServe(":3000", nil))
}