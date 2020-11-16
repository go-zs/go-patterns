package singleton

import "sync"

type Singleton struct {
	name string
}

var (
	once   sync.Once
	single *Singleton
)

func NewSingleton() *Singleton {
	once.Do(func() {
		single = &Singleton{}
	})

	return single
}
