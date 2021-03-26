package cache

import (
	"context"
	"encoding/json"
	"log"
	"rest_api/data"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

var (
	ctx   = context.Background()
	cache = redisCache{
		host:    "localhost:6739",
		db:      0,
		expires: time.Minute * 2,
	}

	client = redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
)

func (cache *redisCache) Set(key string, val *data.Game) {
	json, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}

	client.Set(ctx, key, json, cache.expires*time.Second)

}

func (cache *redisCache) Get(key string) *data.Game {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}

	post := data.Game{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		log.Fatal(err)
	}

	return &post
}
