package utils

import (
	"reflect"
)

func IsErrorValue(ret reflect.Value) (isError bool) {
	// the kind of error object is the `Ptr`, WTF!
	if (ret.Kind() == reflect.Ptr ||
		ret.Kind() == reflect.Interface) &&
		!ret.IsNil() {
		if _, isError = ret.Interface().(error); isError {
			return
		}
	}
	return
}

// CompareFuncs is used to identify whether two functions is equal
func CompareFuncs(func1 interface{}, func2 interface{}) bool {
	// FIXME: comparing two funcs depend on undefined behavor, but it works.
	// Maybe we can find another solution to compare funcs in the future
	sf1 := reflect.ValueOf(func1)
	sf2 := reflect.ValueOf(func2)
	return sf1.Pointer() == sf2.Pointer()
}

// FirstValue return first value of returned value of functions, the function
// should only be invoked in unit test code
func FirstValue(multiples ...interface{}) interface{} {
	// assuming last value is error object.
	if len(multiples) == 0 {
		panic("the number of arguments should be more than 1")
	}

	last := multiples[len(multiples)-1]
	if IsErrorValue(reflect.ValueOf(last)) {
		panic(reflect.ValueOf(last))
	}
	return multiples[0]
}
