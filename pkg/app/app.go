package app

import (
	"fmt"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/transports/http"
	"github.com/gutrse3321/aki/pkg/transports/rpc"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 11:09
 * @Title:
 * --- --- ---
 * @Desc:
 */
/**
从起App传递值获取
*/
type Options struct {
	Name string
	Ip   string
	Port int
	Mode string
}

type Application struct {
	name       string
	logger     *zap.Logger
	opt        *Options
	httpServer *http.Server
	rpcServer  *rpc.Server
}

type Option func(app *Application) error

/**
实例化服务的各自的http、rpc服务器
Option函数参数即是下面的两个HttpServerOption，RpcServerOption
*/
func New(opt *Options, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   opt.Name,
		logger: logger.With(zap.String("type", "Application")),
		opt:    opt,
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

/**
启动微服务的自己http和rpc服务器
*/
func (a *Application) Start() error {
	if a.opt.Port == 0 {
		return errors.New("missing port: 0")
	}

	if a.opt.Ip == "" {
		return errors.New("missing server ip: \"\"")
	}

	addr := fmt.Sprintf("%s:%d", a.opt.Ip, a.opt.Port)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "create tcp serve")
	}

	if a.httpServer != nil {
		if err := a.httpServer.Start(ln); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	if a.rpcServer != nil {
		if err := a.rpcServer.Start(ln); err != nil {
			return errors.Wrap(err, "rpc server start error")
		}
	}

	return nil
}

/**
优雅的关闭http、rpc服务器
*/
func (a *Application) AwaitSignal() {
	//make容量为1的信号
	c := make(chan os.Signal, 1)
	//SIGTERM 直接根据pid杀进程
	//SIGINT  ctrl+c
	//signal.Reset: 重设撤消任何先前呼叫的效果，以通知所提供的信号。如果没有提供信号，所有信号处理程序将被重置。
	//signal.Notify: 将收到的信号，存到channel c中阻塞
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-c:
		//收到信号，执行关闭
		a.logger.Info("receive a signal", zap.String("signal", s.String()))

		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop http server error", zap.Error(err))
			}
		}

		if a.rpcServer != nil {
			if err := a.rpcServer.Stop(); err != nil {
				a.logger.Warn("stop rpc server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}
}

func HttpServerOption(s *http.Server) Option {
	return func(app *Application) error {
		s.ApplicationName(app.name)
		app.httpServer = s
		return nil
	}
}

func RpcServerOption(s *rpc.Server) Option {
	return func(app *Application) error {
		s.ApplicationName(app.name)
		app.rpcServer = s
		return nil
	}
}

var WireSet = wire.NewSet(New)
