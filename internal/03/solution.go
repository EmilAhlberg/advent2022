package main

import (
	"unicode"
)

func SolveP1(rows []string) (total int) {
	for _, row := range rows {
		counts := make(map[rune]int)
		for j, ch := range row {

			if j < (len(row)+1)/2 {
				counts[ch]++
			} else if _, exists := counts[ch]; exists {
				total += getPrio(ch)
				break
			}
		}
	}
	return total
}

func SolveP2(rows []string) (total int) {
	badges := make(map[rune]int)
	for i, row := range rows {
		if i%3 == 0 {
			badges = make(map[rune]int)
		}
		seen := make(map[rune]int)
		for _, ch := range row {
			if _, wasSeen := seen[ch]; !wasSeen {
				badges[ch]++
			}
			seen[ch]++
			if badges[ch] == 3 {
				total += getPrio(ch)
				break
			}
		}
	}
	return total
}

func getPrio(ch rune) int {
	if unicode.IsUpper(ch) {
		return int(ch - 'A' + +27)
	}
	return int(ch - 'a' + 1)

}
