package cache

import (
	"auth-repo/logger"
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func redisOption(redisUrl string, enableTlsMode bool) (*redis.Options, error) {

	opt, err := redis.ParseURL(redisUrl)

	if err != nil {
		tlsConfig := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		opt.TLSConfig = tlsConfig
	}

	return opt, nil
}

func NewRedisClient(redisUrl string, enableRedisTlsMode bool) (*redis.Client, error) {
	opt, err := redisOption(redisUrl, enableRedisTlsMode)
	if err != nil {
		slog.Error("Unable to parse redis url", logger.Extra(map[string]interface{}{
			"err": err.Error(),
		}))
	}

	client := redis.NewClient(opt)
	return client, nil
}
