package main

import (
	"strconv"
)

func Solve(rows []string) (p1, p2 int) {
	M := make([][]int, 0)
	visited := make([][]int, 0)

	for i, row := range rows {
		M = append(M, make([]int, 0))
		visited = append(visited, make([]int, 0))
		for _, ch := range row {
			num, _ := strconv.Atoi(string(ch))
			M[i] = append(M[i], num)
			visited[i] = append(visited[i], 0)
		}
	}

	n, m := len(rows[0]), len(rows)

	var helper func(i, j, max int) int
	helper = func(i, j, max int) int {
		if M[j][i] > max {
			max = M[j][i]
			if visited[j][i] != 2 {
				p1 += 1
				visited[j][i] = 2
			}
		}
		return max
	}

	var helperP2 func(int, int) int
	helperP2 = func(I, J int) int {
		d1, d2, d3, d4 := 0, 0, 0, 0
		for ii := I - 1; ii >= 0; ii-- {
			d1 += 1
			if M[J][ii] >= M[J][I] {
				break
			}
		}

		for ii := I + 1; ii < m; ii++ {
			d2 += 1
			if M[J][ii] >= M[J][I] {
				break
			}
		}
		for jj := J - 1; jj >= 0; jj-- {
			d3 += 1
			if M[jj][I] >= M[J][I] {
				break
			}
		}

		for jj := J + 1; jj < n; jj++ {
			d4 += 1
			if M[jj][I] >= M[J][I] {
				break
			}
		}

		return d1 * d2 * d3 * d4
	}

	for i := 0; i < m; i++ {
		max := -1337
		for j := 0; j < n; j++ {
			max = helper(i, j, max)
			if new := helperP2(i, j); new > p2 {
				p2 = new
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		max := -1337
		for j := n - 1; j > 0; j-- {
			max = helper(i, j, max)
		}
	}
	for j := 0; j < n; j++ {
		max := -1337
		for i := 0; i < m-1; i++ {
			max = helper(i, j, max)
		}
	}
	for j := n - 1; j >= 0; j-- {
		max := -1337
		for i := m - 1; i > 0; i-- {
			max = helper(i, j, max)
		}
	}

	return p1, p2
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
