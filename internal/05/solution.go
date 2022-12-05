package main

import (
	"strconv"
	"strings"
)

func Solve(rows []string, isPart1 bool) string {
	buckets := make([][]string, 0)
	for j := 0; j < 10; j++ {
		buckets = append(buckets, make([]string, 0))
	}
	// Setup
	current := 0
	for i, row := range rows {
		if row == "" {
			current = i
			break
		}
		for j := 0; j*4+1 < len(row); j += 1 {
			if string(row[j*4+1]) == "1" {
				break
			}
			if s := string(row[j*4+1]); s != "" && s != " " {
				buckets[j] = append(buckets[j], string(row[j*4+1]))
			}
		}
	}
	// Crane algorithm
	for i := current; i+1 < len(rows); i++ {
		commands := strings.Split(rows[i+1], " ")
		amount, _ := strconv.Atoi(commands[1])
		source, _ := strconv.Atoi(commands[3])
		target, _ := strconv.Atoi(commands[5])
		temp := append([]string(nil), buckets[source-1][0:amount]...)
		if isPart1 {
			for ii, jj := 0, len(temp)-1; ii < jj; ii, jj = ii+1, jj-1 {
				temp[ii], temp[jj] = temp[jj], temp[ii]
			}
		}
		buckets[target-1] = append(temp, buckets[target-1]...)
		buckets[source-1] = buckets[source-1][amount:]
	}

	//Answer
	var sb strings.Builder
	for i := 0; i < len(buckets); i++ {
		if len(buckets[i]) > 0 {
			sb.WriteString(buckets[i][0])
		}
	}

	return sb.String()
}
