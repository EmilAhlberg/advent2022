package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("04_sample"))
	assert.Equal(t, p1, 2)
	assert.Equal(t, p2, 4)
}

func Test_SolveReal(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("04"))
	assert.Equal(t, p1, 498)
	assert.Equal(t, p2, 859)
}
