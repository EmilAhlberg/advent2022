package main

func Solve(rows []string, windowSize int) int {
	currentWindow, lastSeen := 0, make(map[byte]int)
	for i := 0; i < len(rows[0]); i++ {
		currentWindow++
		ch := rows[0][i]
		if val, wasSeen := lastSeen[ch]; wasSeen && i-val < currentWindow {
			currentWindow = i - val
		}
		if currentWindow == windowSize {
			return i + 1
		}
		lastSeen[ch] = i
	}
	return -1
}
