package tests

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"mux-rest-api/controller"
	"mux-rest-api/entity"
	"mux-rest-api/repository"
	"mux-rest-api/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	ID    int64  = 123
	TITLE string = "Test Title"
	TEXT  string = "Test Text"
)

var (
	postController controller.PostController = controller.NewPostController(postService)
	postService    service.PostService       = service.NewPostService(postRepository)
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
)

func TestAddPost(t *testing.T) {
	//create a new http post request
	var json = []byte(`{"title":"Test Title","text":"Test Text"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(json))

	//Assign Http handler function
	handler := http.HandlerFunc(postController.AddPost)

	//Record http response
	response := httptest.NewRecorder()

	//dispatch http request
	handler.ServeHTTP(response, req)

	//Add assertion on the http status pde amd the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	//decode http response
	var post entity.Post
	json2.NewDecoder(io.Reader(response.Body)).Decode(&post)

	//assert http response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

}

func TestGetPosts(t *testing.T) {
	//insert a new post
	setup()

	//create http request
	req, _ := http.NewRequest("GET", "/posts", nil)

	//Assign Http handler function
	handler := http.HandlerFunc(postController.GetPosts)

	//Record http response
	response := httptest.NewRecorder()

	//dispatch http request
	handler.ServeHTTP(response, req)

	//Add assertion on the http status pde amd the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	//decode http response
	var posts []entity.Post
	json2.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	//assert http response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepository.Save(&post)
}
