package controller

import (
	"encoding/json"
	"mux-rest-api/entity"
	"mux-rest-api/errors"
	"mux-rest-api/service"
	"net/http"
	"strconv"
	"strings"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	GetPostsByID(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}

}
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}
func (*controller) GetPostsByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content Type", "application/json")
	var postId string = strings.Split(request.URL.Path, "/")[2]
	postID, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		panic(err)
	}
	post, err := postService.FindByID(postID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "no posts found"})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling the data"})
		return
	}
	validateErr := postService.Validate(&post)
	if validateErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: validateErr.Error()})
		return
	}
	result, createErr := postService.Create(&post)
	if createErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
