package machine

func (m *Machine) Push(msg interface{}) error {
	return m.entry.Send(msg)
}

func (m *Machine) Subscribe(topic string, size int) chan interface{} {
	return m.middleware.Subscribe(topic, size)
}

func (m *Machine) Unsubscribe(topic string) {
	m.middleware.Unsubscribe(topic)
}

func (m *Machine) Shut() {
	m.cancel()
	for _, v := range m.cancelMap {
		v()
	}
}
