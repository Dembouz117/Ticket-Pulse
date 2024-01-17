package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnectRedisCache(redisAddress string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisAddress)

	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}

// RedisCache implements the Cache interface for Redis.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new instance of RedisCache.
func NewRedisCache(client *redis.Client) *RedisCache { // Return the interface type here
	return &RedisCache{client: client}
}

func (r *RedisCache) SetEntityExpireIn(ctx context.Context, key string, expireInMins int, entity interface{}) error {
	// Serialize the entity to JSON
	serializedEntity, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	// Store the serialized entity in Redis
	err = r.client.Set(ctx, key, serializedEntity, time.Duration(expireInMins)*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) GetEntity(ctx context.Context, key string, entity interface{}) error {
	// Retrieve the serialized entity from Redis
	serializedEntity, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	// Deserialize the entity from JSON
	err = json.Unmarshal([]byte(serializedEntity), entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) DeleteEntity(ctx context.Context, key string) error {
	// Check if the key exists in Redis
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}

	// If the key exists, delete it; otherwise, there's no need to delete
	if exists == 1 {
		err := r.client.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisCache) DeleteEntitiesWithPattern(ctx context.Context, pattern string) error {
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		// Handle error, possibly just log it
		log.Printf("Failed to retrieve keys for cache invalidation: %v", err)
	}

	for _, key := range keys {
		err := r.client.Del(ctx, key)
		if err != nil {
			// Handle caching error, optionally log it
			log.Printf("Failed to invalidate cache for key %s: %v", key, err)
		}
	}
	return nil
}

// BeginTransaction starts a new Redis transaction.
func (r *RedisCache) BeginTransaction(ctx context.Context) redis.Pipeliner {
	return r.client.TxPipeline()
}

// AddToSet adds one or more members to a set stored at key.
func (r *RedisCache) AddToSet(ctx context.Context, key string, members ...interface{}) error {
	_, err := r.client.SAdd(ctx, key, members...).Result()
	return err
}

// GetSetMembers returns all the members of the set stored at key.
func (r *RedisCache) GetSetMembers(ctx context.Context, key string) ([]string, error) {
	members, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}
