package main

import (
	"sort"
	"strconv"
)

func Solve(rows []string) (p1, p2 int) {
	list := []int{}
	mx, current := 0, 0
	for _, row := range rows {
		if row == "" {
			list, current = append(list, current), 0
		} else {
			intVar, _ := strconv.Atoi(row)
			current += intVar
			mx = max(current, mx)
		}
	}
	list = append(list, current)

	sort.Ints(list)
	n := len(list)
	return mx, list[n-1] + list[n-2] + list[n-3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
