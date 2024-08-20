/*
 @Version : 1.0
 @Author  : steven.wong
 @Email   : 'wangxk1991@gamil.com'
 @Time    : 2024/08/20 15:43:10
 Desc     :
*/

package titan

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/piaobeizu/titan/service"
	"github.com/sirupsen/logrus"
)

type Titan struct {
	app       string
	ctx       context.Context
	cancel    context.CancelFunc
	api       *service.ApiServer
	scheduler *service.Scheduler
	singal    chan os.Signal
}

func NewTitan(ctx context.Context, cancel context.CancelFunc, app string) *Titan {
	logrus.Infof("welcome to app %s", app)
	e := &Titan{
		app:    app,
		ctx:    ctx,
		cancel: cancel,
	}
	// we need to register the signals
	signal.Notify(e.singal, syscall.SIGTERM)
	signal.Notify(e.singal, syscall.SIGINT)
	signal.Notify(e.singal, syscall.SIGQUIT)
	signal.Notify(e.singal, syscall.SIGHUP)
	return e
}

func (t *Titan) ApiServer(addr string) *Titan {
	apiCtx := context.WithoutCancel(t.ctx)
	t.api = service.NewApiServer(apiCtx, addr)
	return t
}

func (t *Titan) Scheduler() *Titan {
	schedCtx := context.WithoutCancel(t.ctx)
	t.scheduler = service.NewScheduler(schedCtx)
	return t
}

func (t *Titan) Job(job *service.Job) *Titan {
	if t.scheduler == nil {
		panic("scheduler is nil")
	}
	t.scheduler.Add(job)
	return t
}

func (e *Titan) Start() {
	if e.api != nil {
		e.api.Start()
	}
	if e.scheduler != nil {
		e.scheduler.Start()
	}
	select {
	case <-e.ctx.Done():
		logrus.Info("context done, stop the titan")
		return
	case <-e.singal:
		logrus.Info("receive signal, stop the titan")
		return
	}
}

func (e *Titan) Stop() {
	e.cancel()
	e.api.Stop()
	e.scheduler.Stop()
	logrus.Printf("titan stopped, byebye!")
}