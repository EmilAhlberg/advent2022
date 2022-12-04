package main

import (
	"strconv"
	"strings"
)

func Solve(rows []string) (part1, part2 int) {
	for _, row := range rows {
		ranges := strings.Split(row, ",")
		elf1 := strings.Split(ranges[0], "-")
		elf2 := strings.Split(ranges[1], "-")
		lb1, _ := strconv.Atoi(elf1[0])
		lb2, _ := strconv.Atoi(elf2[0])
		ub1, _ := strconv.Atoi(elf1[1])
		ub2, _ := strconv.Atoi(elf2[1])
		if lb1 <= lb2 && ub1 >= ub2 {
			part1++
		} else if lb2 <= lb1 && ub2 >= ub1 {
			part1++
		}
		if !(lb1 > ub2 || ub1 < lb2) {
			part2 += 1
		}
	}
	return part1, part2
}
