package main

import (
	"fmt"
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
		fmt.Println(I, J, m, n)
		if I == 0 || J == 0 || I == m-1 || J == n-1 {
			return 0
		}

		d1 := 1
		for ii := I - 1; ii > 0; ii-- {
			fmt.Println(M[J][ii], M[J][ii+1])
			if M[J][ii] > M[J][ii+1] {
				break
			}
			d1 += 1
		}

		d2 := 1
		for ii := I + 1; ii < m; ii++ {
			if M[J][ii] > M[J][ii-1] {
				break
			}
			d2 += 1

		}

		d3 := 1
		for jj := J - 1; jj > 0; jj-- {
			if M[jj][I] > M[jj+1][I] {
				break
			}
			d3 += 1
		}

		d4 := 1
		for jj := J + 1; jj < n; jj++ {
			if M[jj][I] > M[jj-1][I] {
				break
			}
			d4 += 1
		}

		fmt.Println(d1, d2, d3, d4)
		return d1 * d2 * d3 * d4
	}
	fmt.Println(M[3][2])
	fmt.Println(M[3])
	helperP2(3, 2)
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
