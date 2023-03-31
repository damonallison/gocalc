package calculator_test

import (
	"fmt"
	"testing"

	"github.com/damonallison/gocalc/calculator"
	"github.com/stretchr/testify/require"
)

func TestAddInts(t *testing.T) {
	fmt.Println("test")
	require.Equal(t, 4, calculator.AddInts(2, 2))
}

func TestAddFloats(t *testing.T) {
	require.InDelta(t, 4.0, calculator.AddFloats(2.0, 2.0), 0.01)
}
