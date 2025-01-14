package cache

import (
	"auth-repo/authentication"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	authentication.Cache
}
type cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &cache{
		client: client,
	}
}
