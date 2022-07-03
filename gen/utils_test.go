package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	seq := []rune{1, 10}
	res := binarySearch(seq, 1)
	assert.Equal(t, 0, res)
}
