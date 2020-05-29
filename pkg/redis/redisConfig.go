package redis

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 16:24
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Options struct {
	Host        string
	Port        int
	Password    string
	Db          int
	Timeout     time.Duration
	ExpiredTime int
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("redis", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal redis config error")
	}

	logger.Info("load redis config success")

	return opt, err
}
