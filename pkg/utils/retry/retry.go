package retry

import (
	"fmt"
	"time"
)

var MaxRetryNumReachError = fmt.Errorf("the max number of retry reached")

type Retry struct {
	//
	RetryConditions []func(r *Retry, err error) bool

	// a series function which will judge error object to indicate the function
	// is executed succesfully. default behavor is err == nil is succeed
	SucceedConditions []func(r *Retry, err error) bool
	// max retry number
	Tries int

	//
	attempts int

	SleepDur time.Duration
}

// Run execute a function repeatedly until the function returned successfully or
// result a result satisfied exit condition
func Run(do func() error, options ...func(r *Retry)) (err error) {
	// TODO
	return nil
}

// Default return a default Retry object which can retry execute a function with
// some speciified retry policy
func Default() *Retry {
	return &Retry{
		Tries:    -1,
		SleepDur: time.Duration(0) * time.Second,
		//
		SucceedConditions: []func(r *Retry, err error) bool{},
		//
		RetryConditions: []func(r *Retry, err error) bool{},
	}
}

// SetOption use variadic functions to set Retry object's option, each function
// could set one or more member of Retry object.
func (r *Retry) SetOption(options ...func(r *Retry)) *Retry {
	for _, opfunc := range options {
		opfunc(r)
	}
	return r
}

// Do function dedicated to executing a function repeatedly until the function
// returned successfully or return a result satisfied exit condition
func (r *Retry) Do(f func() error) error {
	for {
		r.attempts++
		if r.Tries > 0 && r.attempts > r.Tries {
			return MaxRetryNumReachError
		}
		err := f()
		if r.isSucceed(err) {
			return nil
		} else if r.needRetry(err) {
			r.sleepAWhile()
			continue
		} else {
			return err
		}
	}
}

// Attempts return an integer value represented how many time of the function's
// repeat execution
func (r *Retry) Attempts() int {
	return r.attempts
}

func (r *Retry) isSucceed(err error) bool {
	for _, isSucceed := range r.SucceedConditions {
		if isSucceed(r, err) {
			return true
		}
	}

	return err == nil
}

func (r *Retry) needRetry(err error) bool {
	for _, isRetry := range r.RetryConditions {
		if isRetry(r, err) {
			return true
		}
	}
	return false
}

func (r *Retry) sleepAWhile() {
	if r.SleepDur <= 0 {
		return
	}
	time.Sleep(r.SleepDur)
}
