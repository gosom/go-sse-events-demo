package services

import (
	"github.com/go-redis/redis/v8"

	"github.com/gosom/go-sse-events-demo/config"
)

type Container struct {
	Cfg     *config.Config
	Rclient *redis.Client
}

func NewContainer() (*Container, error) {
	var (
		err error
		ans Container
	)
	ans.Cfg, err = config.New()
	if err != nil {
		return nil, err
	}
	ans.Rclient = redis.NewClient(&redis.Options{
		Addr: ans.Cfg.RedisAddr,
	})
	return &ans, nil
}
