package main

func SolveP1(rows []string) (total int) {
	for _, row := range rows {
		total += int(row[2]-'A'-22) + int((row[2]-row[0]+2)%3)*3
	}
	return total
}

func SolveP2(rows []string) (total int) {
	for _, row := range rows {
		total += int(row[2]-'A'-23)*3 + ((int(row[2]-'Y'+3) + int(row[0]-'A')) % 3) + 1
	}
	return total
}

// P2 truth table, second term of expression circular relation
//       A B C
//       0 1 2
//     |--------
// X 2 | 2 0 1
// Y 0 | 0 1 2
// Z 1 | 1 2 0
