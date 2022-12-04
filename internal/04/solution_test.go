package main

import (
	"testing"

	"github.com/EmilAhlberg/advent2022/helpers"
	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("04_sample"))
	assert.Equal(t, p1, 4)
	assert.Equal(t, p2, 2)
}

func Test_SolveReal(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("04"))
	assert.Equal(t, p1, 4)
	assert.Equal(t, p2, 2)
}
