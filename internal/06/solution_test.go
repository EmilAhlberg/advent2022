package main

import (
	"testing"

	helpers "github.com/EmilAhlberg/advent2022/internal/utils"

	"github.com/stretchr/testify/assert"
)

func Test_SolveSample(t *testing.T) {
	windowSizePart1 := 4
	assert.Equal(t, 5, Solve("bvwbjplbgvbhsrlpgdmjqwftvncz", windowSizePart1))
	assert.Equal(t, 6, Solve("nppdvjthqldpwncqszvftbrmjlhg", windowSizePart1))
	assert.Equal(t, 10, Solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", windowSizePart1))
	assert.Equal(t, 11, Solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", windowSizePart1))

	windowSizePart2 := 14
	assert.Equal(t, 19, Solve("mjqjpqmgbljsphdztnvjfqwrcgsmlb", windowSizePart2))
	assert.Equal(t, 23, Solve("bvwbjplbgvbhsrlpgdmjqwftvncz", windowSizePart2))
	assert.Equal(t, 23, Solve("nppdvjthqldpwncqszvftbrmjlhg", windowSizePart2))
	assert.Equal(t, 29, Solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", windowSizePart2))
	assert.Equal(t, 26, Solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", windowSizePart2))

}

func Test_SolveReal(t *testing.T) {
	assert.Equal(t, 1542, Solve(helpers.LoadFile("06")[0], 4))
	assert.Equal(t, 3153, Solve(helpers.LoadFile("06")[0], 14))
}
