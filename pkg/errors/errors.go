package errors

import (
	"fmt"
	"runtime"

	"github.com/juju/errors"
)

const (
	notFoundError = iota
	timeoutError
	conflictError
	badStackStatusError
	badSpecError
	notAllowedError
	unSupport
	notModified
	illegalArgs
	forbidden
	generalError
	breakLoopError
)

type stackError struct {
	code    int
	message string
}

func (e *stackError) Error() string {
	return e.message
}

func (e *stackError) Code() int {
	return e.code
}

// IllegalArgs create an error represent get a invalide argument.
func IllegalArgs(fmts string, messages ...interface{}) error {
	return newStackError(illegalArgs, fmt.Sprintf(fmts, messages...))
}

// NotFound return an error represents something has not been found.
func NotFound(fmts string, messages ...interface{}) error {
	return newStackError(notFoundError, fmt.Sprintf(fmts, messages...))
}

// TimeOut return an error represents do some thing exceed max wait time duration
func TimeOut(fmts string, messages ...interface{}) error {
	return newStackError(timeoutError, fmt.Sprintf(fmts, messages...))
}

// Conflict return an error represents somthing has been conflict
func Conflict(fmts string, messages ...interface{}) error {
	return newStackError(conflictError, fmt.Sprintf(fmts, messages...))
}

// BadStackStatus return an error represents some stack status is invalidate.
func BadStackStatus(fmts string, messages ...interface{}) error {
	return newStackError(badStackStatusError, fmt.Sprintf(fmts, messages...))
}

// BadSpec return an error represents the spec is invalidate.
func BadSpec(fmts string, messages ...interface{}) error {
	return newStackError(badSpecError, fmt.Sprintf(fmts, messages...))
}

// UnSupport return an error represents the required operation is not support.
func UnSupport(fmts string, messages ...interface{}) error {
	return newStackError(unSupport, fmt.Sprintf(fmts, messages...))
}

// NotAllowed return an error represents the required operation is not permitted.
func NotAllowed(fmts string, messages ...interface{}) error {
	return newStackError(notAllowedError, fmt.Sprintf(fmts, messages...))
}

// NotModified return an error represents the required operation does not impact
func NotModified(fmts string, messages ...interface{}) error {
	return newStackError(notModified, fmt.Sprintf(fmts, messages...))
}

// Forbidden return an error represents the required operation are forbidden
func Forbidden(fmts string, messages ...interface{}) error {
	return newStackError(forbidden, fmt.Sprintf(fmts, messages...))
}

// BreakLoopError return an error dedicated to Each function of List object
// indicate loop need to be break
func BreakLoopError(fmts string, message ...interface{}) error {
	return newStackError(breakLoopError, fmt.Sprintf(fmts, message...))
}

// IsNotFound returns true if the specified error was created by NotFound.
func IsNotFound(err error) bool {
	return equalErrorCode(err, notFoundError)
}

// IsTimeOut returns true if the specified error was created by TimeOut.
func IsTimeOut(err error) bool {
	return equalErrorCode(err, timeoutError)
}

// IsConflict determines if the err is an error which indicates the request can
// not be performed because the inner state is conflicts.
func IsConflict(err error) bool {
	return equalErrorCode(err, conflictError)
}

// IsBadStackStatus report whether an error or cause of an error was create with
// BadStackStatus
func IsBadStackStatus(err error) bool {
	return equalErrorCode(err, badStackStatusError)
}

// IsBadSpec report whether an error or cause of an error was create with
// BadSpec
func IsBadSpec(err error) bool {
	return equalErrorCode(err, badSpecError)
}

// IsNotAllowed report whether an error or cause of an error was create with
// BadSpec
func IsNotAllowed(err error) bool {
	return equalErrorCode(err, notAllowedError)
}

// IsUnSupport report whether an error or cause of an error was create with
// UnSupport
func IsUnSupport(err error) bool {
	return equalErrorCode(err, unSupport)
}

// IsNotModified report whether an error or cause of an error was create with
// NotModified
func IsNotModified(err error) bool {
	return equalErrorCode(err, notModified)
}

// IsIllegalArgs report whether an error or cause of an error was create with
// IllegalArgs
func IsIllegalArgs(err error) bool {
	return equalErrorCode(err, illegalArgs)
}

// IsGeneralError report whether an error or cause of an error was create with
// Errorf
func IsGeneralError(err error) bool {
	return equalErrorCode(err, generalError)
}

// IsForbidden report whether an error or cause of an error was created with
// Forbidden
func IsForbidden(err error) bool {
	return equalErrorCode(err, forbidden)
}

// IsBreakLoopError report whether an error or cause of an error was created
// with BreakLoopError
func IsBreakLoopError(err error) bool {
	return equalErrorCode(err, breakLoopError)
}

func equalErrorCode(err error, code int) bool {
	if err == nil {
		return false
	}

	err = errors.Cause(err)
	se, ok := err.(*stackError)
	if ok {
		return se.Code() == code

	}
	return false
}
func newStackError(code int, message string) error {
	err := errors.Annotate(&stackError{
		code, message,
	}, message)
	// we need set location to 2 for getting right linenum and file name
	err.(*errors.Err).SetLocation(2)
	return err
}

// Safely is a helper, which use to emulate try/catch mechanism of java, the
// high-order function of the Safely argument can use panic to return errors and
// make the callee more readable
func Safely(f func() error) func() error {
	// err which is defined in the high-order function declaration is key be-
	// cause the err is the return error object at final.
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					panic(r)
				}
				err = r.(error)
			}
		}()
		ferr := f()
		return ferr
	}
}

// ToErrStack returns a string representation of the annotated error. If the
// error passed as the parameter is not an annotated error, the result is
// simply the result of the Error() method on that error.
// ToErrStack just simply wrapped juju/errors.ToErrStack
func ToErrStack(e error) string {
	return errors.ErrorStack(e)
}

// Cause just simple wrapped juju/errors.Cause function
func Cause(e error) error {
	return errors.Cause(e)
}

// Annotate wrap juju/errors.Annotate function
func Annotate(err error, msg string) error {
	err = errors.Annotate(err, msg)
	if e, ok := err.(*errors.Err); ok {
		// reset filename and line number
		e.SetLocation(1)
	}
	return err
}

// Annotatef wrap juju/errors.Annotatef function
func Annotatef(err error, fmts string, args ...interface{}) error {
	err = errors.Annotatef(err, fmts, args...)
	if e, ok := err.(*errors.Err); ok {
		// reset filename and line number
		e.SetLocation(1)
	}
	return err
}

// Errorf return an error which represent an general case
func Errorf(fmts string, args ...interface{}) error {
	return newStackError(generalError, fmt.Sprintf(fmts, args...))
}

// SetLocation dedicated to change error object file name and file no
func SetLocation(err error) error {
	if e, ok := err.(*errors.Err); ok {
		e.SetLocation(1)
		return e
	}
	return err

}
