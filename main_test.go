package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian(), "Expected median to be 1.5")
	mf.AddNum(3)
	assert.Equal(t, 2.0, mf.FindMedian(), "Expected median to be 2.0")
}