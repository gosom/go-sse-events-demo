package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerAddr string `split_words:"true" default:":8080"`
	Crt        string `default:"/cert/server.crt"`
	Key        string `default:"/cert/server.key"`
	RedisAddr  string `split_words:"true" default:"redis:6379"`
	RedisChan  string `split_words:"true" default:"uuid_chan"`
}

func New() (*Config, error) {
	var cfg Config
	err := envconfig.Process("SSE", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
