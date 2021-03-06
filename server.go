package main

import (
	"fmt"
	"mux-rest-api/cache"
	"mux-rest-api/controller"
	router "mux-rest-api/http"
	"mux-rest-api/repository"
	"mux-rest-api/service"
	"net/http"
	"os"
)

var (
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostsByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))

}
