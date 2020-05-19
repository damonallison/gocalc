package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(4, AddInts(2, 2))
	assert.Equal(2, SubtractInts(4, 2))
	assert.Equal(8, MultiplyInts(4, 2))
	assert.Equal(1, DivideInts(2, 2))
}
