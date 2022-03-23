package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mux-rest-api/entity"
	"mux-rest-api/service"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindByID(id string) (*entity.Post, error) {
	//TODO implement me
	panic("implement me")
}

//Mocking repository Save method for calling on service
func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

//Mocking repository FindAll method for calling on service
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	post := entity.Post{ID: identifier, Title: "Title", Text: "Text"}

	//When
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := service.NewPostService(mockRepo)

	//Expect
	result, _ := testService.FindAll()

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Title", result[0].Title)
	assert.Equal(t, "Text", result[0].Text)

}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	post := entity.Post{ID: identifier, Title: "Title", Text: "Text"}

	//When
	mockRepo.On("Save").Return(&post, nil)

	testService := service.NewPostService(mockRepo)

	result, err := testService.Create(&post)

	//Mock Assertion
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Title", result.Title)
	assert.Equal(t, "Text", result.Text)
	assert.Nil(t, err)

}

//Test for validate function.
func TestValidateEmptyPost(t *testing.T) {
	testService := service.NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "Text"}
	testService := service.NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post title is empty")
}
