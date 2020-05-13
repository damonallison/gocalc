// Package v2 is a simple test calculator created to test go modules
//
// Notice that this is semantic version "v2", which breaks backward compatibility
// with v1. Thus, a new version.
package v2

// AddInts adds 2 ints
func AddInts(x, y int) int {
	return x + y
}
// SubtractInts subtracts 2 ints
func SubtractInts(x, y int) int {
	return x - y
}