package storage

type Infoer interface {
	Info(s string)
}

type noopInfoer struct{}

func newNoopInfoer() *noopInfoer {
	return &noopInfoer{}
}

func (n *noopInfoer) Info(s string) {}
