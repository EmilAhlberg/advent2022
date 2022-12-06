package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	windowSizePart1 := 4
	assert.Equal(t, 5, Solve([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, windowSizePart1))
	assert.Equal(t, 6, Solve([]string{"nppdvjthqldpwncqszvftbrmjlhg"}, windowSizePart1))
	assert.Equal(t, 10, Solve([]string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, windowSizePart1))
	assert.Equal(t, 11, Solve([]string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, windowSizePart1))

	windowSizePart2 := 14
	assert.Equal(t, 19, Solve([]string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, windowSizePart2))
	assert.Equal(t, 23, Solve([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, windowSizePart2))
	assert.Equal(t, 23, Solve([]string{"nppdvjthqldpwncqszvftbrmjlhg"}, windowSizePart2))
	assert.Equal(t, 29, Solve([]string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, windowSizePart2))
	assert.Equal(t, 26, Solve([]string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, windowSizePart2))

}

func Test_SolveReal(t *testing.T) {
	assert.Equal(t, 1542, Solve(helpers.LoadFile("06"), 4))
	assert.Equal(t, 3153, Solve(helpers.LoadFile("06"), 14))

}
