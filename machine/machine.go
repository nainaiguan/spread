package machine

import (
	"context"
	"github.com/nainaiguan/ytchan/api/dft"
	"github.com/nainaiguan/ytchan/api/slow"
	"github.com/nainaiguan/ytchan/api/sub"
	"github.com/nainaiguan/ytchan/dftchan"
	"github.com/nainaiguan/ytchan/slowchan"
	"github.com/nainaiguan/ytchan/subchan"
)

const (
	subTag  = "sub"
	slowTag = "slow"
	DftTag  = "dft"
)

type Machine struct {
	Entry      slow.Slowchan
	Middleware sub.Subchan
	Output     dft.Dftchan
	ctx        context.Context
	cancelMap  cancelMap
	logger     Logger
}

type cancelMap map[string]context.CancelFunc

func DefaultMachine() *Machine {
	m := &Machine{
		ctx:       context.Background(),
		cancelMap: make(map[string]context.CancelFunc),
	}

	m.WithLogger(DefaultLogger())

	sc, cancel := slowchan.Default()
	m.Entry = sc
	m.cancelMap[slowTag] = cancel

	subc, cancel := subchan.Default()
	m.Middleware = subc
	m.cancelMap[subTag] = cancel

	dftc, cancel := dftchan.Default()
	m.Output = dftc
	m.cancelMap[DftTag] = cancel

	return m
}

func NewMachine(cfg *Config) *Machine {
	realCfg := verifyConfig(cfg)
	m := &Machine{
		ctx:       context.Background(),
		cancelMap: make(map[string]context.CancelFunc),
	}
	m.WithLogger(DefaultLogger())

	sc, cancel := slowchan.New(slow.NewSlowArgs{
		Size:           realCfg.EntryCapacity,
		Step:           realCfg.MiddlewareAcceptStep,
		MaxSendProcess: realCfg.MaxEntryPushConcurrency,
	})
	m.Entry = sc
	m.cancelMap[slowTag] = cancel

	subc, cancel := subchan.New(sub.NewSubArgs{
		Size:           realCfg.MiddleCapacity,
		MaxSendProcess: realCfg.MaxMiddlewarePushConcurrency,
	})
	m.Middleware = subc
	m.cancelMap[subTag] = cancel

	dftc, cancel := dftchan.New(dft.NewDftArgs{
		Size:           realCfg.OutputBufferCapacity,
		MaxSendProcess: 1024,
	})
	m.Output = dftc
	m.cancelMap[DftTag] = cancel

	return m
}
