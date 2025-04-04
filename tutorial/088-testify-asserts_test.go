package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	assert.Equal(t, 15, add(6, 9))
	assert.Equal(t, 0, add(1, 2))  // test case fails due to unmet `assert, but test continues because it's an assert
	assert.Equal(t, -5, add(3, 2)) // another unmet `assert, but still continues
	assert.Equal(t, 4, add(2, 2))

	require.Equal(t, 0, add(-3, 4)) //unmet `require` the test will stop immediately

	assert.Equal(t, 6, add(4, 2))
}
