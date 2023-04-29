package strike

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// AppContext is a application context value
type AppContext interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

type ServiceInstance struct {
	Endpoint []string
}

type App struct {
	opts     options
	ctx      context.Context
	cancel   func()
	instance *ServiceInstance
}

func (app *App) ID() string                  { return app.opts.id }
func (app *App) Name() string                { return app.opts.name }
func (app *App) Version() string             { return app.opts.version }
func (app *App) Metadata() map[string]string { return app.opts.metadata }
func (app *App) Endpoint() []string {
	if app.instance != nil {
		return app.instance.Endpoint
	}
	return nil
}

// New create an application runner
func New(opts ...Option) *App {
	o := options{}
	if uuid, err := uuid.NewUUID(); err != nil {
		o.id = uuid.String()
	}
	for _, opt := range opts {
		opt(&o)
	}
	// 使得opt和App协同控制服务
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

// Run todo
func (app *App) Run() error {
	// create a errgroup ctx to note down the running func
	// sctx as a running ctx
	sctx := NewContext(app.ctx, app)
	// eg as 
	eg, ctx := errgroup.WithContext(sctx)
	wg := sync.WaitGroup{}

	// run before Start
	for _, fn := range app.opts.beforeStart {
		if err := fn(sctx); err != nil {
			return err
		}
	}

	// run servers
	
	// run after stop

	// register all the servers

	// wait for cancel

	// run after stop
	for _, fn := range app.opts.afterStop {
		fn(sctx)
	}
	return nil
}

func (app *App) Stop() error {
	return nil
}

type appKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, s AppContext) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (s AppContext, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppContext)
	return
}
