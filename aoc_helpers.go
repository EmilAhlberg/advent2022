package helpers

import (
	"bufio"
	"fmt"
	"os"
)

const verbose = false

func LoadFile(name string) []string {
	file, err := os.Open(fmt.Sprintf("../input/%s.txt", name))
	defer file.Close()

	var rows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	if err != nil {
		fmt.Println("yippie ki-yay", name)
		os.Exit(1)
	}

	if verbose {
		for _, row := range rows {
			fmt.Println(row)
		}
	}

	return rows
}
