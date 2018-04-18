package utils

import (
	"sync"
	"time"
)

// WaitTimeout will wait on WatiGroup timeout seconds, if it return true value
// which indicate that wait operation is timedout
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
