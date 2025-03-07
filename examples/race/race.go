// Package race is an example to demonstrate race conditions.
// We use `go test -race` to check if there are any race conditions.
// But the race detector only works when your tests cover the code you want to check.
package race

import (
	"sync"
	"time"
)

var (
	m = make(map[string]int)
	l = sync.RWMutex{}
)

func ConcurrentOpMap() {
	go func() {
		for i := 0; i < 100; i++ {
			getFromMap("key")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			setMap("key", i)
		}
	}()
	time.Sleep(time.Second * 2)
}

func getFromMap(k string) int {
	l.RLock()
	defer l.RUnlock()
	return m[k]
}

func setMap(k string, v int) {
	l.Lock()
	defer l.Unlock()
	m[k] = v
}
