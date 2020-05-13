// Package calculator is a simple desk calculator created to test go modules
package calculator

// AddInts adds integers
func AddInts(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// AddFloats adds floats
func AddFloats(values ...float64) float64 {
	total := 0.0
	for _, num := range values {
		total += num
	}
	return total
}
