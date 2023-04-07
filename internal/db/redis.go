package db

import (
	"context"

	"github.com/redis/go-redis/v9"

	errs "github.com/3JoB/ulib/err"
)

var (
	rd  *redis.Client
	ctx = context.Background()
)

func init() {
	rd = redis.NewClient(&redis.Options{
		Addr:     kc.String("cache.addr"),
		Password: kc.String("cache.pwd"), // no password set
		DB:       kc.Int("cache.db"),     // use default DB
	})
	if err := rd.Conn().Ping(ctx).Err(); err != nil {
		err := &errs.Err{Op: "internal/db/redis.Ping", E: err}
		panic(err)
	}
}

func NewRedis() *redis.Client {
	return rd
}
