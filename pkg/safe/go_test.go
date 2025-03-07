package safe_test

import (
	"bytes"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/safe"
)

type logger struct {
	sync.Mutex
	buffer bytes.Buffer
}

func (l *logger) Error(msg string, vv ...any) {
	l.Lock()
	defer l.Unlock()
	l.buffer.WriteString(msg)
	for _, v := range vv {
		l.buffer.WriteString(": ")
		l.buffer.WriteString(cast.ToString(v))
	}
}

func TestSafeGo(t *testing.T) {
	num := atomic.Int64{}
	l := &logger{}
	safe.Callback(func(_ any, _ []byte) {
		num.Add(1)
	})
	safe.CallCallBack(func(err any, stack []byte) {
		l.Error("safe call occurs panic", err, string(stack))
	})
	safe.GoCallBack(func(err any, stack []byte) {
		l.Error("safe go occurs panic", err, string(stack))
	})

	safe.Go(func() {
		panic("panic in safe.Go")
	}, func(_ any, _ []byte) {
		num.Add(1)
	})

	start := time.Now().UnixMilli()
	safe.Call(func() {
		time.Sleep(10 * time.Millisecond)
		panic("panic in safe.Call")
	}, func(_ any, _ []byte) {
		num.Add(1)
	})
	safe.Wait()
	end := time.Now().UnixMilli()

	assert.Equal(t, int64(4), num.Load())
	assert.True(t, end-start >= 10)

	buffer := l.buffer.String()
	assert.Contains(t, buffer, "safe call occurs panic: panic in safe.Call")
	assert.Contains(t, buffer, "safe go occurs panic: panic in safe.Go")
}
