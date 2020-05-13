package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInts(t *testing.T) {
	assert.Equal(t, 4, AddInts(2, 2))
}
