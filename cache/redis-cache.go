package cache

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"mux-rest-api/entity"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

//Create a new Redis Client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.Post) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}
func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	return &post
}
