package service

import (
	"errors"
	"math/rand"
	"mux-rest-api/entity"
	"mux-rest-api/repository"
	"strconv"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(postRepo repository.PostRepository) PostService {
	repo = postRepo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*service) FindByID(id string) (*entity.Post, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}
