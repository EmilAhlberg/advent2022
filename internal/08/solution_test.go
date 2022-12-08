package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("08_sample"))
	assert.Equal(t, 21, p1)
	assert.Equal(t, 8, p2)
}

func Test_SolveReal(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("08"))
	assert.Equal(t, 1662, p1)
	assert.Equal(t, 537600, p2)
}
