package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr     string
	Password string
}

func (c *Config) GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
	})
	log.Info("create redis client done")
	return client
}
