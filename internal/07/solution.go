package main

import (
	"strconv"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func Solve(rows []string) (total, deleteSize int) {
	currentPath := ""
	directories := make(map[string]int)
	for _, row := range rows {
		commandline := strings.Split(row, " ")
		if len(row) > 0 && row[0] == '$' {
			if commandline[1] == "cd" {
				if commandline[2] == "/" {
					currentPath = "/"
				} else if commandline[2] == ".." {
					temp := strings.Split(currentPath, "/")
					if len(temp) == 1 {
						currentPath = "/"
					} else {
						currentPath = strings.Join(temp[:len(temp)-1], "/")
					}
				} else {
					if len(currentPath) != 1 {
						currentPath = currentPath + "/"
					}
					currentPath = currentPath + commandline[2]
				}
			}
		} else if commandline[0] != "dir" {
			cheat := strings.Split(currentPath, "/")
			for i := 0; i < len(cheat); i++ {
				size, _ := strconv.Atoi(commandline[0])
				paths := strings.Join(cheat[:i+1], "/")
				directories[paths] += size
			}
		}

	}

	rootFolderTrick := 0
	for k, v := range directories {
		if path := strings.Split(k, "/"); len(path) == 2 && path[1] != "" {
			rootFolderTrick += v
		}
	}

	directories["/"] += rootFolderTrick

	var _ []kv

	total, deleteSize = 0, 999999999990
	p2Target := 30000000 - (70000000 - directories["/"])
	for _, val := range directories {
		if val <= 100000 {
			total += val
		}
		if val >= p2Target && val-p2Target < deleteSize {
			deleteSize = val
		}
	}

	return total, deleteSize
}
