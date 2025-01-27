package cache

import (
	"auth-repo/logger"
	"crypto/tls"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func redisOption(redisUrl string, enableTlsMode bool) (*redis.Options, error) {

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	if enableTlsMode {
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
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
