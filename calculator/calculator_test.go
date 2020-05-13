package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInts(t *testing.T) {
	assert.Equal(t, 6, AddInts(1, 2, 3))
}
func TestAddFloats(t *testing.T) {
	assert.Equal(t, 6.0, AddFloats(1.0, 2.0, 3.0))
}
