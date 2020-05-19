// Package calculator is a simple test calculator created to test go modules
//
// Notice that this is semantic version "v2", which breaks backward compatibility
// with v1. Thus, a new version.
package calculator

// AddInts adds 2 ints
func AddInts(x, y int) int {
	return x + y
}

// SubtractInts subtracts 2 ints
func SubtractInts(x, y int) int {
	return x - y
}

// MultiplyInts multiplies 2 ints
func MultiplyInts(x, y int) int {
	return x * y
}

// DivideInts divides 2 ints
func DivideInts(x, y int) int {
	return x / y
}
