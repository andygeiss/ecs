package ecs

import "runtime"

func init() {
	runtime.LockOSThread()
}

// Run simplifies the engine usage by calling the Setup(), Run() and Teardown() internally.
func Run(em *EntityManager, sm *SystemManager) {
	engine := NewEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Run()
}

var fnMain = func(fn func()) {}

// RunAsMain runs the given function in the main OS thread.
// This is necessary for non-Go library functions that depend on per-thread state.
func RunAsMain(fn func()) {
	fnMain(fn)
}

// Prepare prepares Go for running Cgo calls in a separate worker function safely by
// locking the main OS thread to the current Goroutine.
func Prepare(worker func()) {
	callQueue := make(chan func())
	fnMain = func(fn func()) {
		done := make(chan bool, 1)
		callQueue <- func() {
			fn()
			done <- true
		}
		<-done
	}
	// Spawn a new Goroutine for non-sensitive workers.
	go func() {
		worker()
		close(callQueue)
	}()
	for fn := range callQueue {
		fn()
	}
}
