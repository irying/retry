package utils

import (
	"reflect"
)

// ArrayClone use reflect implement generic array clone function, why not use
// copy? because  we need deep copy element in a slice, if element of a slice
// contains Clone  method, the ArrayClone would recursively invoke Clone
// function of the element of the slice.
// Caution: reflect might bring performance penalty , use it carefully
func ArrayClone(arr interface{}) (dest interface{}) {
	arrValue := reflect.ValueOf(arr)

	// Retrieve the type, and check if it is one of the array or slice.
	arrType := arrValue.Type()
	arrElemType := arrType.Elem()
	if arrType.Kind() != reflect.Array && arrType.Kind() != reflect.Slice {
		panic("Array parameter's type is neither array nor slice.")

	}

	resultSliceType := reflect.SliceOf(arrElemType)
	resultSlice := reflect.MakeSlice(resultSliceType, 0, arrValue.Len())
	for i := 0; i < arrValue.Len(); i++ {
		arrElemAddrType := arrValue.Index(i).Addr().Type()
		m, ok := arrElemType.MethodByName("Clone")
		if ok {
			r := m.Func.Call([]reflect.Value{arrValue.Index(i)})
			resultSlice = reflect.Append(resultSlice, r[0])
		} else if m, ok := arrElemAddrType.MethodByName("Clone"); ok {
			r := m.Func.Call([]reflect.Value{arrValue.Index(i).Addr()})
			resultSlice = reflect.Append(resultSlice, r[0])
		} else {
			resultSlice = reflect.Append(resultSlice, arrValue.Index(i))
		}

	}
	return resultSlice.Interface()
}

// DiffStrSlice return a slice which contains the difference of two string slice
func DiffStrSlice(slice1, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break

				}

			}

			if !found {
				diff = append(diff, s1)

			}

		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1

		}

	}

	return diff
}

// StrInSlice return a bool value to report whether a string be contained in a
// string slice
func StrInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true

		}
	}
	return false
}
