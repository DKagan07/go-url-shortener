package store

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

const CacheDuration = time.Hour * 6

var ctx = context.Background()

func InitializeStore() *StorageService {
	r := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	status, err := r.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing redis: %+v\n", err))
	}
	fmt.Println("Redis started successfully, pong message: ", status)

	return &StorageService{
		redisClient: r,
	}
}

func (s *StorageService) SaveUrlMapping(shortUrl, originalUrl, userId string) {
	_, err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Result()
	if err != nil {
		panic(fmt.Sprintf("SaveUrlMapping::redisClient.Set::%+v", err))
	}
}

func (s *StorageService) RetreiveInitialUrl(shortUrl string) string {
	origUrl, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("RetrieveInitialUrl::redisClient.Get::%+v", err))
	}
	return origUrl
}
