package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_02_2(t *testing.T) {
	rows := LoadFile("aoc_02")
	total := 0
	for _, row := range rows {
		total += int(row[2]-'A'-22) + int((row[2]-row[0]+2)%3)*3

	}
	fmt.Println(total)
	assert.False(t, true)

}

func Test_02(t *testing.T) {
	rows := LoadFile("aoc_02")

	sum := 0
	sum2 := 0
	for _, row := range rows {

		sum += score(row[0], row[2])
		sum2 += score2(row[0], row[2])
	}
	fmt.Println(sum)
	fmt.Println(sum2)
	assert.False(t, true)
}

func score(they, you byte) int {
	total := 0

	//if you == 'Y' {
	//	total += 2
	//} else if you == 'X' {
	//	total += 1
	//} else if you == 'Z' {
	//	total += 3
	//}

	if you == 'X' {
		if they == 'B' {
			return total
		}
		if they == 'C' {
			return total + 6
		}
	}

	if you == 'Y' {
		if they == 'A' {
			return total + 6
		}
		if they == 'C' {
			return total
		}
	}

	if you == 'Z' {
		if they == 'A' {
			return total
		}
		if they == 'B' {
			return total + 6
		}
	}

	return total + 3
}

func score2(they, you byte) int {
	total := 0

	if you == 'X' {
		total += 0
	} else if you == 'Y' {
		total += 3
	} else if you == 'Z' {
		total += 6
	}

	if you == 'X' {
		if they == 'A' {
			return total + 3
		}
		if they == 'B' {
			return total + 1
		}
		if they == 'C' {
			return total + 2
		}
	}

	if you == 'Y' {
		if they == 'A' {
			return total + 1
		}
		if they == 'B' {
			return total + 2
		}
		if they == 'C' {
			return total + 3
		}
	}

	if you == 'Z' {
		if they == 'A' {
			return total + 2
		}
		if they == 'B' {
			return total + 3
		}
		if they == 'C' {
			return total + 1
		}
	}

	return total + 3
}
