package main

func Solve(rows string, windowSize int) int {
	currentWindow, lastSeenAt := 0, make(map[rune]int)
	for i, ch := range rows {
		currentWindow++
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
