package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	assert.Equal(t, 15, SolveP1(helpers.LoadFile("02_sample")))
	assert.Equal(t, 12, SolveP2(helpers.LoadFile("02_sample")))
}

func Test_SolveReal(t *testing.T) {
	assert.Equal(t, 14069, SolveP1(helpers.LoadFile("02")))
	assert.Equal(t, 12411, SolveP2(helpers.LoadFile("02")))
}
