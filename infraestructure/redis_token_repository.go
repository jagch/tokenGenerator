package infraestructure

import (
	"context"
	"jagch/tokenGenerator/domain"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisTokenRepository struct {
	redisClient *redis.Client
}

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return redisClient
}

func NewRedisTokenRepository(redisClient *redis.Client) domain.TokenRepository {
	return &redisTokenRepository{
		redisClient: redisClient,
	}
}

func (r *redisTokenRepository) CreateTokens(key string, value any) error {
	ctx := context.Background()
	switch os.Getenv("ENVIRONMENT") {
	case "dev":
		err := r.redisClient.Set(ctx, key, value, 60*time.Minute).Err()
		if err != nil {
			return err
		}
	case "prod":
		err := r.redisClient.Set(ctx, key, value, 24*time.Hour).Err()
		if err != nil {
			return err
		}
	default:
		err := r.redisClient.Set(ctx, key, value, 60*time.Minute).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *redisTokenRepository) TokenExists(ctx context.Context, key string) bool {
	exists, _ := r.redisClient.Exists(ctx, key).Result()
	return exists == 1
}
