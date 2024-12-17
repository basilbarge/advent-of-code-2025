package main

import (
	util "github.com/aoc2024/utils"
	"fmt"
)

func main() {
	horizontalLines := util.ReadLines("../sample.txt")

	verticalLines := make([]string, len(horizontalLines[0]))

	for _, line := range horizontalLines {

		for idx, char := range line {
			verticalLines[idx] += string(char)
		}
		fmt.Println(line)
	}

	for _, vertLine := range verticalLines {
		fmt.Println(vertLine)
	}
}
