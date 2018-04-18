package utils

// Int32P return a pointer value of a int32
func Int32P(i int32) *int32 {
	return &i
}

// IntP return a pointer value of a int
func IntP(i int) *int {
	return &i
}

// FalseP return a pointer value of a bool false value
func FalseP() *bool {
	f := false
	return &f
}

// TrueP return a pointer value of a bool true value
func TrueP() *bool {
	f := true
	return &f
}
