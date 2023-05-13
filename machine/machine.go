package machine

import (
	"context"
	"github.com/nainaiguan/ytchan/api/slow"
	"github.com/nainaiguan/ytchan/api/sub"
	"github.com/nainaiguan/ytchan/slowchan"
	"github.com/nainaiguan/ytchan/subchan"
)

const (
	subTag  = "sub"
	slowTag = "slow"
)

type Machine struct {
	entry      slow.Slowchan
	middleware sub.Subchan
	ctx        context.Context
	cancel     context.CancelFunc
	cancelMap  cancelMap
	logger     Logger
}

type cancelMap map[string]context.CancelFunc

func DefaultMachine(f func(interface{}) interface{}) *Machine {
	ctx, cancel := context.WithCancel(context.Background())
	m := &Machine{
		ctx:       ctx,
		cancel:    cancel,
		cancelMap: make(map[string]context.CancelFunc),
	}

	m.WithLogger(DefaultLogger())

	sc, cancel := slowchan.Default()
	m.entry = sc
	m.cancelMap[slowTag] = cancel

	subc, cancel := subchan.Default()
	m.middleware = subc
	m.cancelMap[subTag] = cancel

	go m.daemon(m.ctx, f)

	return m
}

func NewMachine(cfg *Config, f func(interface{}) interface{}) *Machine {
	realCfg := verifyConfig(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	m := &Machine{
		ctx:       ctx,
		cancel:    cancel,
		cancelMap: make(map[string]context.CancelFunc),
	}
	m.WithLogger(DefaultLogger())

	sc, cancel := slowchan.New(slow.NewSlowArgs{
		Size:           realCfg.EntryCapacity,
		Step:           realCfg.MiddlewareAcceptStep,
		MaxSendProcess: realCfg.MaxEntryPushConcurrency,
	})
	m.entry = sc
	m.cancelMap[slowTag] = cancel

	subc, cancel := subchan.New(sub.NewSubArgs{
		Size:           realCfg.MiddleCapacity,
		MaxSendProcess: realCfg.MaxMiddlewarePushConcurrency,
	})
	m.middleware = subc
	m.cancelMap[subTag] = cancel

	go m.daemon(m.ctx, f)

	return m
}
