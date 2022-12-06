package main

func Solve(input string, windowSize int) int {
	currentWindow, lastSeenAt := 0, make(map[rune]int)
	for i, ch := range input {
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
