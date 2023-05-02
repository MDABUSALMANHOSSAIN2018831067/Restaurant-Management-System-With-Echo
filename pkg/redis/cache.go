package redis

import (
	"context"
	"encoding/json"
	"restaurant-management/pkg/models"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore() *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisStore{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rs *RedisStore) Set(ID string, user []models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rs.client.Set(rs.ctx, ID, data, 0).Err()
}

func (rs *RedisStore) Get(id string) (*[]models.User, error) {
	var msg []models.User
	data, err := rs.client.Get(rs.ctx, id).Bytes()

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
