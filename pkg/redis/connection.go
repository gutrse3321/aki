package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 16:30
 * @Title:
 * --- --- ---
 * @Desc:
 */

func New(opt *Options) (*redis.Client, error) {
	redisOpt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		Password:     opt.Password,
		DB:           opt.Db,
		ReadTimeout:  opt.Timeout * time.Second,
		WriteTimeout: opt.Timeout * time.Second,
	}

	client := redis.NewClient(redisOpt)
	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.Wrap(err, "redis create client error")
	}

	return client, nil
}
