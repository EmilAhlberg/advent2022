package aoc

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func Test_03_1(t *testing.T) {
	rows := LoadFile("aoc_03")

	total := 0
	for _, row := range rows {

		counts := make(map[rune]int)
		for j, ch := range row {

			if j < (len(row)+1)/2 {
				counts[ch]++
			} else if _, exists := counts[ch]; exists {

				if unicode.IsUpper(ch) {
					fmt.Println(string(ch), int(ch-'A'+27))
					total += int(ch - 'A' + +27)
				} else {
					fmt.Println(string(ch), int(ch-'a'+1))
					total += int(ch - 'a' + 1)

				}
				break
			}
		}
	}

	fmt.Println("total", total)
	assert.False(t, true)
}

func Test_03_2(t *testing.T) {
	rows := LoadFile("aoc_03")

	total := 0
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

			if val, _ := badges[ch]; val == 3 {
				if unicode.IsUpper(ch) {
					fmt.Println(string(ch), int(ch-'A'+27))
					total += int(ch - 'A' + +27)
				} else {
					fmt.Println(string(ch), int(ch-'a'+1))
					total += int(ch - 'a' + 1)

				}
				break
			}
		}
	}

	fmt.Println("total", total)
	assert.False(t, true)
}
