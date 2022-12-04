package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	assert.Equal(t, 157, SolveP1(helpers.LoadFile("03_sample")))
	assert.Equal(t, 70, SolveP2(helpers.LoadFile("03_sample")))
}

func Test_SolveReal(t *testing.T) {
	assert.Equal(t, 8072, SolveP1(helpers.LoadFile("03")))
	assert.Equal(t, 2567, SolveP2(helpers.LoadFile("03")))
}
