package main

import (
	"fmt"
	"regexp"

	util "github.com/aoc2024/utils"
)

func main() {
	horizontalLines := util.ReadLines("../sample.txt")
	verticalLines := make([]string, len(horizontalLines[0]))

	numDiagonals := (len(horizontalLines) + len(horizontalLines[0]) - 1)

	forwardDiagonal := make([]string, numDiagonals)
	//backwardDiagonal := make([]string, numDiagonals)

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

	for idx := 0; idx < len(horizontalLines)+len(verticalLines)-1; idx++ {
		xcount := 0
		if idx > len(verticalLines) {
			xcount = idx - len(verticalLines)
		}
		ycount := min(idx, len(horizontalLines)-1)

		forwardDiagonal[idx] = string(horizontalLines[ycount][xcount])

		xcount += 1
		ycount -= 1

		for (xcount < len(verticalLines)) && (ycount >= 0) {
			forwardDiagonal[idx] += string(horizontalLines[ycount][xcount])
			xcount++
			ycount--
		}
	}

	fmt.Println("Forward Diagonal")
	for _, forDiag := range forwardDiagonal {
		fmt.Println(forDiag)
	}
}
