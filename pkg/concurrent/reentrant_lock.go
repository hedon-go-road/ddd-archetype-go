package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// ReentrantLock is a reentrant lock support multiple locks in the same goroutine
type ReentrantLock struct {
	sync.Mutex

	// owner is the id of the lock holder goroutine
	owner atomic.Int64
	// reentrant is the number of reentrant lock
	reentrant atomic.Int64
}

func (m *ReentrantLock) Lock() {
	gid := goid.Get()

	if m.owner.Load() == gid {
		m.reentrant.Add(1)
		return
	}
	m.Mutex.Lock()
	m.owner.Store(gid)
	m.reentrant.Store(1)
}

func (m *ReentrantLock) Unlock() {
	gid := goid.Get()
	if m.owner.Load() != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d", m.owner.Load(), gid))
	}
	if m.reentrant.Add(-1) != 0 {
		return
	}
	m.owner.Store(-1)
	m.Mutex.Unlock()
}
