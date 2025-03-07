package safe

import (
	"runtime/debug"
	"sync"
)

// sg is a zconfig instance of safeGo which holds the panic callbacks and a sync.WaitGroup.
// Each program only has one instance of safeGo.
var sg = safeGo{
	goCallbacks:   make([]PanicCallback, 0),
	callCallbacks: make([]PanicCallback, 0),
}

// safeGo struct contains a WaitGroup and slices of panic callbacks.
type safeGo struct {
	sync.WaitGroup

	// goCallbacks are the callbacks that will be invoked when a panic occurs in safe.Go.
	goCallbacks []PanicCallback

	// callCallbacks are the callbacks that will be invoked when a panic occurs in safe.Call.
	callCallbacks []PanicCallback
}

// PanicCallback defines a function type that takes an error and a stack trace as parameters.
type PanicCallback = func(err any, stack []byte)

// Callback appends the provided callbacks to both safe.Go and safe.Call.
func Callback(callbacks ...PanicCallback) {
	sg.goCallbacks = append(sg.goCallbacks, callbacks...)
	sg.callCallbacks = append(sg.callCallbacks, callbacks...)
}

// GoCallBack appends the zconfig callbacks for safe.Go,
// it should be used only in the main package.
func GoCallBack(callbacks ...PanicCallback) {
	sg.goCallbacks = append(sg.goCallbacks, callbacks...)
}

// CallCallBack appends the zconfig callbacks for safe.Call,
// it should be used only in the main package.
func CallCallBack(callbacks ...PanicCallback) {
	sg.callCallbacks = append(sg.callCallbacks, callbacks...)
}

// Call safely executes the provided function f and recovers from any panic,
// invoking the registered panic callbacks with the error and stack trace.
func Call(f func(), panicCallback ...PanicCallback) {
	defer func() {
		debug.Stack()
		if err := recover(); err != nil {
			stack := debug.Stack()
			for _, callback := range sg.callCallbacks {
				callback(err, stack)
			}
			for _, callback := range panicCallback {
				callback(err, stack)
			}
		}
	}()
	f()
}

// Go safely executes the provided function f in a new goroutine and recovers from any panic,
// invoking the registered panic callbacks with the error and stack trace.
// NOTE: don't use it in the function never return!!!!!
func Go(f func(), panicCallback ...PanicCallback) {
	sg.Add(1)
	go func() {
		defer sg.Done()
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				for _, callback := range sg.goCallbacks {
					callback(err, stack)
				}
				for _, callback := range panicCallback {
					callback(err, stack)
				}
			}
		}()
		f()
	}()
}

// Wait blocks until the WaitGroup counter is zero, meaning all goroutines have finished.
func Wait() {
	sg.Wait()
}
