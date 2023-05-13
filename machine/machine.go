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
		logger:    DefaultLogger(),
	}

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
