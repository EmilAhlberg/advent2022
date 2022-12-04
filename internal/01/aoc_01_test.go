package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnInvalid(t *testing.T) {
	rows := LoadFile("01")

	list := []int{}
	mx := 0
	current := 0
	for _, row := range rows {
		if row == "" {
			list = append(list, current)
			current = 0
		} else {
			intVar, err := strconv.Atoi(row)
			fmt.Println(err)
			current += intVar
			mx = max(current, mx)
		}
	}

	fmt.Println("ANSWER:", mx)
	sort.Ints(list)
	fmt.Println(list)

	sum := 0

	for i := len(list) - 1; i > len(list)-4; i-- {
		sum += list[i]
		fmt.Println(sum)
	}

	assert.False(t, true)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
