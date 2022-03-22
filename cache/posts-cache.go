package cache

import "mux-rest-api/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
