package machine

import (
	"context"
)

func (m *Machine) daemon(ctx context.Context, f func(msg interface{}) interface{}) {
	entryChan := m.entry.Start(8192)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-entryChan:
				if err := m.middleware.Send(f(msg)); err != nil {
					m.logger.Errorf("push to middle err: %s", err)
					continue
				}
			}
		}
	}()
}
