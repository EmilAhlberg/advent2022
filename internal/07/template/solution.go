package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve(rows []string) (ans int) {

	currentPath := ""
	directories := make(map[string]int)

	currentDir := ""
	for _, row := range rows {
		fmt.Println(currentPath)
		if len(row) > 0 && row[0] == '$' {
			line := strings.Split(row, " ")
			command := line[1]
			fmt.Print(line)
			if command == "cd" {
				if line[2] == "/" {
					currentPath = "/"
				} else if line[2] == ".." {
					fmt.Println("hej")
					temp := strings.Split(currentPath, "/")
					fmt.Println(temp, len(temp))
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

		} else if temp := strings.Split(row, " "); temp[0] == "dir" {
			currentDir = temp[1]
			fmt.Println(currentDir)
		} else {
			temp := strings.Split(row, " ")[0]
			size, _ := strconv.Atoi(temp)
			cheat := strings.Split(currentPath, "/")

			for i := 0; i < len(cheat); i++ {

				te := strings.Join(cheat[:i+1], "/")
				if te == "" {
					continue
				}
				//fmt.Println(te, "Aa")
				//fmt.Println(te, size)
				directories[te] += size
			}
		}

	}
	//delete(directories, "/")
	fmt.Print(directories["/"])
	mytemp := 0
	for k, v := range directories {
		fmt.Println(len(strings.Split(k, "/")), strings.Split(k, "/"))
		if path := strings.Split(k, "/"); len(path) == 2 && path[1] != "" {
			mytemp += v
		}
	}

	directories["/"] += mytemp

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range directories {
		ss = append(ss, kv{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	total := 0
	// Print the x top values

	// 4 HERE FOR TEST
	for _, kv := range ss[:4] {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		//cheat := strings.Split(kv.Key, "/")
		if len(kv.Key) == 1 {
			total += kv.Value
		}
	}
	fmt.Println(directories)

	total = 0
	// 4 HERE FOR TEST
	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		//cheat := strings.Split(kv.Key, "/")
		if kv.Value <= 100000 {
			total += kv.Value
		}
	}

	target := 30000000 - (70000000 - directories["/"])
	fmt.Println("target", target)
	currentSelection := 999999999990

	for _, kv := range ss {
		if kv.Value < target {
			continue
		} else if kv.Value-target < currentSelection {
			currentSelection = kv.Value
			fmt.Println(currentSelection)
		}
	}
	fmt.Print(directories["/"], 70000000, 70000000-directories["/"])

	return total
}
