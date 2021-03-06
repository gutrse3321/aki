package http

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 11:11
 * @Title: 公共http服务器
 * --- --- ---
 * @Desc:
 */

/**
从配置文件中获取 key: http
*/
type Options struct {
	Name string
	Mode string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	opt := &Options{}

	if err := v.UnmarshalKey("app", opt); err != nil {
		return nil, err
	}

	return opt, nil
}

type InitControllers func(r *gin.RouterGroup)

type Middleware func(r *gin.Engine)

/**
初始化http服务器
*/
func NewRouter(opt *Options, logger *zap.Logger, init InitControllers, mid Middleware) *gin.Engine {
	gin.SetMode(opt.Mode)

	r := gin.New()

	/**
	中间件使用
	*/
	//panic自动恢复，无需自行重启
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	mid(r)

	init(r.Group(opt.Name))

	return r
}

/**
http服务器结构体
*/
type Server struct {
	opt        *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

/**
注入依赖项，返回实例
*/
func New(opt *Options, logger *zap.Logger, router *gin.Engine) (*Server, error) {
	s := &Server{
		opt:    opt,
		logger: logger,
		router: router,
	}

	return s, nil
}

/**
设置http服务的名称
*/
func (s *Server) ApplicationName(name string) {
	s.app = name
}

/**
启动http服务器
*/
func (s *Server) Start(ln net.Listener) error {
	//配置标准包http server配置, 指定处理器为gin
	s.httpServer = http.Server{
		Handler: s.router,
	}

	s.logger.Info("http server starting ...")

	//启动服务器，监听端口，以及异常处理
	go func() {
		if err := s.httpServer.Serve(ln); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server: ", zap.String("application", s.app))

	//等待5秒后再继续
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

var WireSet = wire.NewSet(New, NewRouter, NewOptions)
