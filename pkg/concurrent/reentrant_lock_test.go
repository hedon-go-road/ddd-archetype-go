package concurrent

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReentrantLock(t *testing.T) {
	rl := new(ReentrantLock)

	// unlock without lock should panic
	assert.Panics(t, rl.Unlock)

	// unlock other goroutines lock should panic
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		rl.Lock()
		time.Sleep(2 * time.Millisecond) // wait for other goroutine unlock
		rl.Unlock()
	}()
	assert.Panics(t, func() {
		time.Sleep(1 * time.Millisecond) // make sure the other goroutine lock
		rl.Unlock()
	})
	wg.Wait()

	// lock and unlock multiple times in one goroutine should success
	rl.Lock()
	rl.Lock()
	assert.Equal(t, int64(2), rl.reentrant.Load())
	rl.Unlock()
	rl.Unlock()

	// unlock out of reentrant count should panic
	assert.Panics(t, rl.Unlock)
}
