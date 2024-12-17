package main

import (
	"fmt"
	"regexp"

	util "github.com/aoc2024/utils"
)

func main() {
	horizontalLines := util.ReadLines("../sample.txt")

	verticalLines := make([]string, len(horizontalLines[0]))

	regexp.MustCompile("XMAS|SAMX")

	fmt.Println("Horizontal")
	for _, line := range horizontalLines {

		for idx, char := range line {
			verticalLines[idx] += string(char)
		}
		fmt.Println(line)
	}

	fmt.Println("Vertical")
	for _, vertLine := range verticalLines {
		fmt.Println(vertLine)
	}
}
