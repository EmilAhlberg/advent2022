package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	ans := Solve(helpers.LoadFile("07_sample"))
	assert.Equal(t, ans, 1)
}

func Test_SolveReal(t *testing.T) {
	ans := Solve(helpers.LoadFile("07"))
	assert.Equal(t, ans, 1)

}
