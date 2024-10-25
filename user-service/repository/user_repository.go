package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/pranay999000/microservices-grpc/user-service/domain"
)

type UserRepo struct {
	UserMap 	*map[int32]*domain.User
}

func NewUserRepo(userMap *map[int32]*domain.User) *UserRepo {
	return &UserRepo{
		UserMap: userMap,
	}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *domain.User) error {

	var err error

	if _, ok := (*repo.UserMap)[user.Id]; ok {
		log.Printf("id %d already exists\n", user.Id)
		err = errors.New("user id " + strconv.Itoa(int(user.Id)) + " already exists")
		return err
	}

	(*repo.UserMap)[user.Id] = user
	return nil

}

func (repo *UserRepo) GetUserById(ctx context.Context, userId int32) (*domain.User, error) {

	var err error

	if _, ok := (*repo.UserMap)[userId]; !ok {
		log.Printf("id %d does not exists\n", userId)
		err = errors.New("user id " + strconv.Itoa(int(userId)) + " does not exists")
		return nil, err
	}

	return (*repo.UserMap)[userId], nil

}