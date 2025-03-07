package collection

import (
	"sync"
)

type Manager[K comparable, T any] struct {
	sync.RWMutex
	items map[K]T
}

func New[K comparable, T any]() *Manager[K, T] {
	return &Manager[K, T]{
		items: make(map[K]T, 1024),
	}
}

func (m *Manager[K, T]) Get(id K) T {
	m.RLock()
	defer m.RUnlock()
	return m.items[id]
}

func (m *Manager[K, T]) Add(id K, item T) {
	m.Lock()
	defer m.Unlock()
	m.items[id] = item
}

func (m *Manager[K, T]) Delete(id K) T {
	m.Lock()
	defer m.Unlock()
	item := m.items[id]
	delete(m.items, id)
	return item
}

func (m *Manager[K, T]) Exists(id K) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[id]
	return ok
}

func (m *Manager[K, T]) Range(f func(K, T) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, t := range m.items {
		if !f(k, t) {
			break
		}
	}
}
