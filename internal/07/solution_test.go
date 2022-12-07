package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("07_sample"))
	assert.Equal(t, 95437, p1)
	assert.Equal(t, 24933642, p2)
}

func Test_SolveReal(t *testing.T) {
	p1, p2 := Solve(helpers.LoadFile("07"))
	assert.Equal(t, 1543140, p1)
	assert.Equal(t, 1117448, p2)

}
