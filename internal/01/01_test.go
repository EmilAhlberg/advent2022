package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("01_sample"))
	assert.Equal(t, 24000, p1)
	assert.Equal(t, 45000, p2)
}

func Test_SolveReal(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("01"))
	assert.Equal(t, 67658, p1)
	assert.Equal(t, 200158, p2)
}
