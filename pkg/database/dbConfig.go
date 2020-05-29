package database

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 16:37
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Options struct {
	User         string
	Password     string
	Host         string
	Db           string
	MaxIdleConns int
	MaxOpenConns int
	Debug        bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("database", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal database config error")
	}

	logger.Info("load database config success", zap.String("host", opt.Host))

	return opt, err
}
