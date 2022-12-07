package main

import (
	"sort"
	"strconv"
	"strings"
)

func Solve(rows []string) (total, deleteSize int) {
	currentPath := ""
	directories := make(map[string]int)
	for _, row := range rows {
		if len(row) > 0 && row[0] == '$' {
			line := strings.Split(row, " ")
			command := line[1]
			if command == "cd" {
				if line[2] == "/" {
					currentPath = "/"
				} else if line[2] == ".." {
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
					currentPath = currentPath + line[2]
				}
			}
		} else if paths := strings.Split(row, " "); paths[0] != "dir" {
			cheat := strings.Split(currentPath, "/")
			for i := 0; i < len(cheat); i++ {
				size, _ := strconv.Atoi(strings.Split(row, " ")[0])
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

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range directories {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	total = 0

	for _, kv := range ss[:4] {
		if len(kv.Key) == 1 {
			total += kv.Value
		}
	}

	total = 0
	for _, kv := range ss {
		if kv.Value <= 100000 {
			total += kv.Value
		}
	}

	target := 30000000 - (70000000 - directories["/"])
	deleteSize = 999999999990

	for _, kv := range ss {
		if kv.Value < target {
			continue
		} else if kv.Value-target < deleteSize {
			deleteSize = kv.Value
		}
	}

	return total, deleteSize
}
