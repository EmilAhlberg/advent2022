package main

func Solve(rows string, windowSize int) int {
	currentWindow, lastSeenAt := 0, make(map[byte]int)
	for i := 0; i < len(rows); i++ {
		ch, currentWindow := rows[i], currentWindow+1
		if idx, wasSeen := lastSeenAt[ch]; wasSeen && i-idx < currentWindow {
			currentWindow = i - idx
		}
		if currentWindow == windowSize {
			return i + 1
		}
		lastSeenAt[ch] = i
	}
	return -1
}
