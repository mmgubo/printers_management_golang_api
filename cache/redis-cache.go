package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"mfundo.com/printers/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PrinterCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, printer *entity.Printer) {
	client := cache.getClient()

	// serialize Post object to JSON
	json, err := json.Marshal(printer)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}
func (cache *redisCache) Get(key string) *entity.Printer {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	printer := entity.Printer{}
	err = json.Unmarshal([]byte(val), &printer)
	if err != nil {
		panic(err)
	}

	return &printer
}


