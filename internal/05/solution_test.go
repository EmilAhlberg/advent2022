package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	assert.Equal(t, Solve(helpers.LoadFile("05_sample"), true), "CMZ")
	assert.Equal(t, Solve(helpers.LoadFile("05_sample"), false), "MCD")
}

func Test_SolveReal(t *testing.T) {
	assert.Equal(t, Solve(helpers.LoadFile("05"), true), "VRWBSFZWM")
	assert.Equal(t, Solve(helpers.LoadFile("05"), false), "RBTWJWMCF")
}
