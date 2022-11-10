package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/idprm/go-alesse/src/pkg/util/localconfig"
)

func NewRedisClient() *redis.Client {

	secret, err := localconfig.LoadSecret("src/server/secret.yaml")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     secret.RD.Host,
		Password: secret.RD.Password,
		DB:       0,
	})
	return client
}
